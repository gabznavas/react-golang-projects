package controllers

import (
	todoUsecases "api/usecases/todos"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAllTodosController struct {
	usecase *todoUsecases.GetAllTodosUsecase
}

func NewGetAllTodosController(usecase *todoUsecases.GetAllTodosUsecase) *GetAllTodosController {
	return &GetAllTodosController{usecase: usecase}
}

func (c *GetAllTodosController) GetAllTodos(ctx *gin.Context) {
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid offset"})
		return
	}
	if offset < 0 {
		ctx.JSON(400, gin.H{"error": "Offset must be greater than 0"})
		return
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid limit"})
		return
	}
	if limit < 0 {
		ctx.JSON(400, gin.H{"error": "Limit must be greater than 0"})
		return
	}
	if limit > 50 {
		ctx.JSON(400, gin.H{"error": "Limit must be less than 50"})
		return
	}

	todoResponses, err := c.usecase.GetAllTodos(ctx, offset, limit)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, todoResponses)
}
