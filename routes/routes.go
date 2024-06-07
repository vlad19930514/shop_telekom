package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vlad19930514/shop_telekom/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/user/username", UpdateName)
}
