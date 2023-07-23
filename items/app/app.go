package app

import (
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	*/
	deps := BuildDependencies()
	MapUrls(router, deps)
	_ = router.Run(":8090")
}
