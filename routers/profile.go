package routers

import (
	"ewallet_be/controllers"
	"ewallet_be/middlewares"

	"github.com/gin-gonic/gin"
)

func editProfileRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.PATCH("", controllers.EditUser)
	r.POST("/picture", controllers.UploadProfilePicture)
}

func listUsers(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("", controllers.ListUsersForTransfer)
}
