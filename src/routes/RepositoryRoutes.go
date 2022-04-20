package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucchesisp/go-test-egroupas/src/usecases/gitRepo"
)

func addRepositoriesRoute(rg *gin.RouterGroup) {
	repositories := rg.Group("/repositories")

	repositories.GET("/", gitRepo.HandleController)
}
