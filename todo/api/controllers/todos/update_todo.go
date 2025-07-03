package controllers

import (
	dto "api/dto"
	todoUsecases "api/usecases/todos"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateTodoController struct {
	usecase *todoUsecases.UpdateTodoUsecase
}

func NewUpdateTodoController(usecase *todoUsecases.UpdateTodoUsecase) *UpdateTodoController {
	return &UpdateTodoController{usecase: usecase}
}

func (c *UpdateTodoController) UpdateTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	var updateTodoRequest dto.UpdateTodoRequest
	if err := ctx.ShouldBindJSON(&updateTodoRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = c.usecase.UpdateTodo(ctx, id, updateTodoRequest)
	if err != nil {
		if err == todoUsecases.ErrTodoNotFound {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if err == todoUsecases.ErrTodoAlreadyExists {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, nil)
}
