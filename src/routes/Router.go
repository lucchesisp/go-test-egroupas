package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucchesisp/go-test-egroupas/src/config"
)

func Run(port string) {
	ginMode := config.GetEnvVariable("GIN_MODE")

	gin.SetMode(ginMode)

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	GetRoutes(router)

	router.Run(":" + port)
}

func GetRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	addRepositoriesRoute(v1)
}
