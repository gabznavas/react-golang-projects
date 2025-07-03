package controllers

import (
	todoUsecases "api/usecases/todos"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetTodoController struct {
	usecase *todoUsecases.GetTodoUsecase
}

func NewGetTodoController(usecase *todoUsecases.GetTodoUsecase) *GetTodoController {
	return &GetTodoController{usecase: usecase}
}

func (c *GetTodoController) GetTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	todoResponse, err := c.usecase.GetTodo(ctx, id)
	if err != nil {
		if err == todoUsecases.ErrTodoNotFound {
			ctx.JSON(404, gin.H{"error": "Todo not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, todoResponse)
}
