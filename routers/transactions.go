package routers

import (
	"ewallet_be/controllers"

	"github.com/gin-gonic/gin"
)

func topupRouter(r *gin.RouterGroup) {
	r.POST("", controllers.Topup)
}
