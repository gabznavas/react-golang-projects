package controllers

import (
	projectUsecase "api/usecases/project"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProjectByIdController struct {
	getProjectByIdUsecase *projectUsecase.GetProjectByIdUsecase
}

func NewGetProjectByIdController(getProjectByIdUsecase *projectUsecase.GetProjectByIdUsecase) *GetProjectByIdController {
	return &GetProjectByIdController{getProjectByIdUsecase: getProjectByIdUsecase}
}

func (c *GetProjectByIdController) GetProjectById(ctx *gin.Context) {
	projectId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid project ID"})
		return
	}
	project, err := c.getProjectByIdUsecase.Execute(uint(projectId))
	if err != nil {
		if errors.Is(err, projectUsecase.ErrProjectNotFound) {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, project)
}
