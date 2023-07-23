package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func mapUrls(router *gin.Engine, dependencies *Dependencies) {
	// mapeo de usuarios
	router.GET("/user/:id", dependencies.UserController.GetUserById)
	router.POST("/user", dependencies.UserController.InsertUser)
	router.DELETE("/user/:id", dependencies.UserController.DeleteUser)
	router.POST("/login", dependencies.UserController.Login)

	log.Info("Terminando mapeo de urls")
}
