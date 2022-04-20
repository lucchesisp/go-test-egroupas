package gitRepo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleController(context *gin.Context) {

	lastProjects, err := HandleService.lastProjects(5)

	if err != nil {
		context.JSON(http.StatusInternalServerError, ResponseBodyDTO{
			Error: err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, ResponseBodyDTO{
		LastProjects: lastProjects,
	})
}
