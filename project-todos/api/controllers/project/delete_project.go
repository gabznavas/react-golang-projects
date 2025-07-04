package controllers

import (
	projectUsecase "api/usecases/project"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProjectController struct {
	deleteProjectUsecase *projectUsecase.DeleteProjectUsecase
}

func NewDeleteProjectController(deleteProjectUsecase *projectUsecase.DeleteProjectUsecase) *DeleteProjectController {
	return &DeleteProjectController{deleteProjectUsecase: deleteProjectUsecase}
}

func (c *DeleteProjectController) DeleteProject(ctx *gin.Context) {
	projectId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid project ID"})
		return
	}

	err = c.deleteProjectUsecase.Execute(uint(projectId))
	if err != nil {
		if errors.Is(err, projectUsecase.ErrProjectNotFound) {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(204, nil)
}
