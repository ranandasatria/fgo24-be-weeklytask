package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Missing Authorization header"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("APP_SECRET")), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
			ctx.Abort()
			return
		}

		userIdFloat, ok := claims["userId"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "userId not found in token"})
			ctx.Abort()
			return
		}

		ctx.Set("userId", int(userIdFloat))
		ctx.Next()
	}
}
