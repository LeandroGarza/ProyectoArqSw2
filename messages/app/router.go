package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func mapUrls(router *gin.Engine, dependencies *Dependencies) {
	// middleware que valida la request
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/messages/:id", dependencies.Controller.GetMessageById)
	router.GET("/messages/item/:itemid", dependencies.Controller.GetMessagesByItem)
	router.DELETE("/messages/user/:userid", dependencies.Controller.DeleteMessagesByUser)

	// middleware que valida el token de usuario
	router.Use(dependencies.Controller.ValidateToken)

	router.POST("/messages", dependencies.Controller.CreateMessage)
	router.GET("/messages/user/:userid", dependencies.Controller.GetMessageByUser)
	router.DELETE("/messages/:id", dependencies.Controller.DeleteMessage)
}
