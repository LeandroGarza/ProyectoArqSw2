package app

import (
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	deps := BuildDependencies()
	mapUrls(router, deps)
	_ = router.Run(":9001")
}
