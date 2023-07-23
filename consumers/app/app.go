package app

import (
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	BuildDependencies()
	_ = router.Run(":9003")
}
