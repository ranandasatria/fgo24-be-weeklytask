package controllers

import (
	"ewallet_be/models"
	"ewallet_be/utils"
	"net/http"
	"os"
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

	userID, err := models.Register(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "User created",
		Results: user,
	})

	if err := models.CreateWalletForUser(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "User created but failed to create wallet",
		})
		return
	}
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

	if form.PIN != user.PIN {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Wrong PIN",
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
	var input models.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	userIdRaw, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	userID := userIdRaw.(int)

	oldUser, err := models.FindOneUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	if input.Email == "" {
		input.Email = oldUser.Email
	}
	if input.Password == "" {
		input.Password = oldUser.Password
	}
	if input.PIN == "" {
		input.PIN = oldUser.PIN
	}
	if input.Username == "" {
		input.Username = oldUser.Username
	}
	if input.Phone == "" {
		input.Phone = oldUser.Phone
	}
	if input.ProfilePicture == "" {
		input.ProfilePicture = oldUser.ProfilePicture
	}

	if err := models.EditUser(userID, input); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User updated",
		Results: input,
	})
}
