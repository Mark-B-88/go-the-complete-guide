package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // Configures HTTP Server

	server.GET("/api/v1/events", getEvents)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}