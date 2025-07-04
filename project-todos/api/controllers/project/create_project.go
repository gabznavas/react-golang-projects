package controllers

import (
	"api/dto"
	projectUsecase "api/usecases/project"

	"github.com/gin-gonic/gin"
)

type CreateProjectController struct {
	createProjectUsecase *projectUsecase.CreateProjectUsecase
}

func NewCreateProjectController(createProjectUsecase *projectUsecase.CreateProjectUsecase) *CreateProjectController {
	return &CreateProjectController{createProjectUsecase: createProjectUsecase}
}

func (c *CreateProjectController) CreateProject(ctx *gin.Context) {
	params := &dto.CreateProjectRequest{}
	if err := ctx.ShouldBindBodyWithJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	project, err := c.createProjectUsecase.Execute(params)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, project)
}
