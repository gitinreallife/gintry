package controller

import (
	"modules/models"

	"github.com/gin-gonic/gin"
)

func FindLASTUsers(c *gin.Context) {
	// var lasTUser []models.LASTUser
	models.ReadLASTUsers()
}
