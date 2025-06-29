package controllers

import (
	"ewallet_be/models"
	"ewallet_be/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	if err := models.Register(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: true,
			Message: "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User created",
		Results: user,
	})
}

func Login(ctx *gin.Context) {
	godotenv.Load()
	secretKey := os.Getenv("APP_SECRET")

	form := struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		PIN      string `json:"pin" binding:"required"`
	}{}

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	user, err := models.FindOneUserByEmail(form.Email)

	if err != nil {
		//handle
	}

	if user == (models.User{}) || (form.Password != user.Password) {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Wrong email or password",
		})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	token, _ := generateToken.SignedString([]byte(secretKey))

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Login success",
		Results: map[string]string{
			"token": token,
		},
	})
}

func EditUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	userID, _ := strconv.Atoi(id)
	if err := models.EditUser(userID, user); err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, utils.Response{
				Success: false,
				Message: "User ID not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User updated",
		Results: user,
	})
}
