package routers

import (
	"ewallet_be/controllers"
	"ewallet_be/middlewares"

	"github.com/gin-gonic/gin"
)

func editProfileRouter(r *gin.RouterGroup){
	r.Use(middlewares.VerifyToken())
	r.PATCH(":id", controllers.EditUser)
}