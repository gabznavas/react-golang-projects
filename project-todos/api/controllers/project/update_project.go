package controllers

import (
	"api/dto"
	projectUsecase "api/usecases/project"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProjectController struct {
	updateProjectUsecase *projectUsecase.UpdateProjectUsecase
}

func NewUpdateProjectController(updateProjectUsecase *projectUsecase.UpdateProjectUsecase) *UpdateProjectController {
	return &UpdateProjectController{updateProjectUsecase: updateProjectUsecase}
}

func (c *UpdateProjectController) UpdateProject(ctx *gin.Context) {
	projectId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid project ID"})
		return
	}
	var params dto.UpdateProjectRequest
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = c.updateProjectUsecase.Execute(uint(projectId), params)
	if err != nil {
		if errors.Is(err, projectUsecase.ErrProjectNotFound) {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, projectUsecase.ErrProjectAlreadyExists) {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(204, nil)
}
