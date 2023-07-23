package app

import "github.com/gin-gonic/gin"

func mapUrls(router *gin.Engine, dependencies *Dependencies) {
	router.GET("/user/:id", dependencies.Cachecontroller.GetUserData)
}
