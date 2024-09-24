package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", home)
	app.Run(":3000")
}

func home(response *gin.Context) {
	response.JSON(http.StatusOK, gin.H{"message": "chinna pilalu"})
}
