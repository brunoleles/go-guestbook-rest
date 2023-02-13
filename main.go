package main

import (
	"net/http"

	"main/controllers"
	"main/support"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	server := gin.Default()

	// Ping test
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	server.GET("/guestbook", controllers.GuestbookGet)
	server.POST("/guestbook", controllers.GuestbookPost)
	server.DELETE("/guestbook", controllers.GuestbookDelete)

	return server
}

func main() {
	support.InitEnv()
	support.InitDB()

	server := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	server.Run(":8080")
}
