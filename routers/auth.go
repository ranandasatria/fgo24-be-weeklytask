package routers

import (
	"ewallet_be/controllers"

	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.RouterGroup) {
	r.POST("", controllers.Register)
}
