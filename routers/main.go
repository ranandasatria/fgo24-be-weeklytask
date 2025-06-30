package routers

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine) {
	registerRouter(r.Group("/register"))
	loginRouter(r.Group("/login"))
	editProfileRouter(r.Group("/profile"))
	listUsers(r.Group("/users"))
	walletRouter(r.Group("/wallet"))
	topupRouter(r.Group("/topup"))
	transferRouter(r.Group("/transfer"))
}