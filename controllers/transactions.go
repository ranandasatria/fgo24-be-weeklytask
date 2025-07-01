package controllers

import (
	"context"
	"encoding/json"
	"ewallet_be/models"
	"ewallet_be/utils"
	"net/http"
	"strings"

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

func TransferHistory(ctx *gin.Context) {
	userIdRaw, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{Success: false, Message: "Unauthorized"})
		return
	}

	idUser := userIdRaw.(int)
	keyword := ctx.Query("keyword")

	data, err := models.GetTransferHistory(idUser, keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to fetch transfer history",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Transfer history fetched",
		Results: data,
	})
}

func ListUsersForTransfer(ctx *gin.Context) {
	err := utils.RedisClient().Ping(context.Background()).Err()
	noredis := false
	if err != nil && strings.Contains(err.Error(), "refused") {
		noredis = true
	}

	if !noredis {
		result := utils.RedisClient().Exists(context.Background(), ctx.Request.RequestURI)
		if result.Val() != 0 {
			users := []models.UserListItem{}
			data := utils.RedisClient().Get(context.Background(), ctx.Request.RequestURI)
			if err := json.Unmarshal([]byte(data.Val()), &users); err == nil {
				ctx.JSON(http.StatusOK, utils.Response{
					Success: true,
					Message: "List users (from Redis)",
					Results: users,
				})
			}
		}
	}


	userIdRaw, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	idUser := userIdRaw.(int)
	keyword := ctx.Query("keyword")

	users, err := models.GetOtherUsers(idUser, keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to fetch users",
			Errors:  err.Error(),
		})
		return
	}

	if !noredis {
		encoded, err := json.Marshal(users)
		if err == nil {
			utils.RedisClient().Set(context.Background(), ctx.Request.RequestURI, string(encoded), 0)
		}
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User list fetched",
		Results: users,
	})
}
