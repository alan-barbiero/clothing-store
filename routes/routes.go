package routes

import (
	"clothing-store/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/clothing", controllers.CreateClothing)
	r.GET("/clothing", controllers.GetAllClothing)
	r.PUT("/clothing/:id", controllers.UpdateClothing)
	r.DELETE("/clothing/:id", controllers.DeleteClothing)

	return r
}
