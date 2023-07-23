package app

import (
	"github.com/gin-gonic/gin"
)

// MapUrls maps the urls
func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	router.GET("/search=:searchQuery", dependencies.SearchController.Search)
	router.GET("/search/byuser=:userid", dependencies.SearchController.SearchByUserId)
	router.POST("/", dependencies.SearchController.InsertItems)
	router.DELETE("/all", dependencies.SearchController.DeleteAll)
	router.DELETE("/:userid", dependencies.SearchController.DeleteByUserId)
}
