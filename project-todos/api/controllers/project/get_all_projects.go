package controllers

import (
	projectUsecase "api/usecases/project"

	"github.com/gin-gonic/gin"
)

type GetAllProjectsController struct {
	getAllProjectsUsecase *projectUsecase.GetAllProjectsUsecase
}

func NewGetAllProjectsController(getAllProjectsUsecase *projectUsecase.GetAllProjectsUsecase) *GetAllProjectsController {
	return &GetAllProjectsController{getAllProjectsUsecase: getAllProjectsUsecase}
}

func (c *GetAllProjectsController) GetAllProjects(ctx *gin.Context) {
	projects, err := c.getAllProjectsUsecase.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, projects)
}
