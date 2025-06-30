package routers

import (
	"ewallet_be/controllers"
	"ewallet_be/middlewares"

	"github.com/gin-gonic/gin"
)

func walletRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("", controllers.GetWallet)
}
