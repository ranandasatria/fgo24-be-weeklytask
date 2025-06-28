package controllers

import (
	"ewallet_be/models"
	"ewallet_be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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
