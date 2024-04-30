package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vlad19930514/shop_telekom/db"
	"github.com/vlad19930514/shop_telekom/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
