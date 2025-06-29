package routers

import (
	"ewallet_be/controllers"
	"ewallet_be/middlewares"

	"github.com/gin-gonic/gin"
)

func topupRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.POST("", controllers.Topup)
}

func transferRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.POST("", controllers.Transfer)
}
