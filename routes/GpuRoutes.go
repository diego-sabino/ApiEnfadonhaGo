package routes

import (
	"mymodule/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/gpu", controllers.AllGpus)
	r.POST("/gpu", controllers.CreateNewGpu)
	r.GET("/gpu/:id", controllers.GetGpuById)
	r.DELETE("/gpu/:id", controllers.DeleteGpu)
	r.Run()
}
