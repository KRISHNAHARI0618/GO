package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("", getEvents)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"hello": "world dear"})

}
