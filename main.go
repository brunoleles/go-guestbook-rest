package main

import (
	"net/http"

	"main/controllers"
	. "main/database"
	"main/models"
	"main/support"

	"github.com/gin-gonic/gin"
)

func newServer() *gin.Engine {
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

func doMigrations() {
	support.PanicOnError(DB().AutoMigrate(&models.GuestbookModel{}))
}

func main() {
	support.InitEnv()
	InitDB()
	doMigrations()

	server := newServer()

	// Listen and Server in 0.0.0.0:8080
	server.Run(":8080")
}
