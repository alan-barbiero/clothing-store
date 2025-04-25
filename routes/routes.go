package routes

import (
	"clothing-store/controllers"

	"github.com/gin-gonic/gin"
)

func ClothingRoutes(r *gin.RouterGroup) {
	r.POST("/clothing", controllers.CreateClothing)
	r.GET("/clothing", controllers.GetAllClothing)
	r.PUT("/clothing/:id", controllers.UpdateClothing)
	r.DELETE("/clothing/:id", controllers.DeleteClothing)
}

func ShoeRoutes(r *gin.RouterGroup) {
	r.POST("/shoe", controllers.CreateShoe)
	r.GET("/shoe", controllers.GetAllShoes)
	r.PUT("/shoe/:id", controllers.UpdateShoe)
	r.DELETE("/shoe/:id", controllers.DeleteShoe)
}
