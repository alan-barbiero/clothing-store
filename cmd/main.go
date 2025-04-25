package main

import (
	"clothing-store/config"
	"clothing-store/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		routes.ClothingRoutes(api)
		routes.ShoeRoutes(api)
	}

	r.Static("/assets", "./frontend/dist/assets")

	fmt.Println("🚀 Server on http://localhost:6688")
	err := http.ListenAndServe(":6688", r)
	if err != nil {
		log.Fatal(err)
	}
}
