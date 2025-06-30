package controllers

import (
	"ewallet_be/models"
	"ewallet_be/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWallet(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}
	userID := userIdRaw.(int)

	wallet, err := models.GetWalletByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to fetch wallet",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Wallet fetched",
		Results: wallet,
	})
}
