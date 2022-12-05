package controllers

import (
	"mymodule/database"
	"mymodule/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllGpus(c *gin.Context) {
	var Gpu []models.Gpu
	database.DB.Find(&Gpu)
	c.JSON(http.StatusOK, Gpu)
}

func DeleteGpu(c *gin.Context) {
	id := c.Params.ByName("id")
	var Gpu []models.Gpu
	database.DB.Delete(&Gpu, id)
	c.JSON(http.StatusNoContent, Gpu)
}

func CreateNewGpu(c *gin.Context) {
	var Gpu models.Gpu
	if err := c.ShouldBindJSON(&Gpu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateGpuData(&Gpu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&Gpu)
	c.JSON(http.StatusCreated, Gpu)
}
