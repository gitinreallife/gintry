package main

import (
	"modules/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectLASDB()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	// referenced from https://levelup.gitconnected.com/build-your-first-rest-api-in-go-language-using-gin-framework-827aadc14e07
	r.Run(":3000")
}
