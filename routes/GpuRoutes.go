package routes

import (
	"mymodule/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/gpu", controllers.AllGpus)
	r.POST("/gpu", controllers.CreateNewGpu)
	r.DELETE("/gpu/:id", controllers.DeleteGpu)
	r.Run()
}
