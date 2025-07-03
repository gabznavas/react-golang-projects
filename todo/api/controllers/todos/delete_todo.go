package controllers

import (
	todoUsecases "api/usecases/todos"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteTodoController struct {
	usecase *todoUsecases.DeleteTodoUsecase
}

func NewDeleteTodoController(usecase *todoUsecases.DeleteTodoUsecase) *DeleteTodoController {
	return &DeleteTodoController{usecase: usecase}
}

func (c *DeleteTodoController) DeleteTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	err = c.usecase.DeleteTodo(ctx, id)
	if err != nil {
		if err == todoUsecases.ErrTodoNotFound {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, nil)
}
