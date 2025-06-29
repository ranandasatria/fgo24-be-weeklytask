package controllers

import (
	"ewallet_be/models"
	"ewallet_be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Topup(ctx *gin.Context) {
	var topup models.Topup

	if err := ctx.ShouldBind(&topup); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
			Errors:  err.Error(),
		})
		return
	}

	if err := models.CreateTopup(topup); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to top up",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Topup successful",
		Results: topup,
	})
}


func Transfer(ctx *gin.Context) {
	var transfer models.Transfer

	if err := ctx.ShouldBind(&transfer); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
			Errors:  err.Error(),
		})
		return
	}

	if err := models.CreateTransfer(transfer); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Transfer failed",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Transfer successful",
		Results: transfer,
	})
}