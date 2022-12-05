package controllers

import (
	"mymodule/database"
	"mymodule/models"

	"github.com/gin-gonic/gin"
)

func AllGpus(c *gin.Context) {
	var Gpu []models.Gpu
	database.DB.Find(&Gpu)
	c.JSON(200, Gpu)
}
