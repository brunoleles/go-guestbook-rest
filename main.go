package main

import (
	"net/http"

	. "main/database"
	"main/models"
	"main/modules/guestbook"
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

	server.GET("/guestbook", guestbook.GuestbookGet)
	server.POST("/guestbook", guestbook.GuestbookPost)
	server.DELETE("/guestbook", guestbook.GuestbookDelete)

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
