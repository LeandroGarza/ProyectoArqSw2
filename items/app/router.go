package app

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {

	router.GET("/items/:id", dependencies.ItemController.Get)
	router.DELETE("/items/user/:userid", dependencies.ItemController.DeleteByUserId)

	// Middlewares para validar solicitud y token
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(dependencies.ItemController.ValidateToken)
	router.POST("/items", dependencies.ItemController.InsertItems)

	fmt.Println("Finishing mappings configurations")
}
