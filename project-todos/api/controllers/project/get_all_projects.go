package controllers

import (
	projectUsecase "api/usecases/project"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAllProjectsController struct {
	getAllProjectsUsecase *projectUsecase.GetAllProjectsUsecase
}

func NewGetAllProjectsController(getAllProjectsUsecase *projectUsecase.GetAllProjectsUsecase) *GetAllProjectsController {
	return &GetAllProjectsController{getAllProjectsUsecase: getAllProjectsUsecase}
}

func (c *GetAllProjectsController) GetAllProjects(ctx *gin.Context) {
	search := ctx.Query("search")
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	projects, err := c.getAllProjectsUsecase.Execute(search, offset, limit)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, projects)
}
