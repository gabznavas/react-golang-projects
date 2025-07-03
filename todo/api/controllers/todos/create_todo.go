package controllers

import (
	todoUsecases "api/usecases/todos"

	dto "api/dto"

	"github.com/gin-gonic/gin"
)

type CreateTodoController struct {
	usecase *todoUsecases.CreateTodoUsecase
}

func NewCreateTodoController(usecase *todoUsecases.CreateTodoUsecase) *CreateTodoController {
	return &CreateTodoController{usecase: usecase}
}

func (c *CreateTodoController) CreateTodo(ctx *gin.Context) {
	var createTodoRequest dto.CreateTodoRequest
	if err := ctx.ShouldBindJSON(&createTodoRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todoResponse, err := c.usecase.CreateTodo(ctx, createTodoRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, todoResponse)
}
