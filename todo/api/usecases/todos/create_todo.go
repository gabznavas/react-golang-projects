package usecases

import (
	dto "api/dto"
	models "api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateTodoUsecase struct {
	db *gorm.DB
}

func NewCreateTodoUsecase(db *gorm.DB) *CreateTodoUsecase {
	return &CreateTodoUsecase{
		db: db,
	}
}

func (u *CreateTodoUsecase) CreateTodo(ctx *gin.Context, createTodoRequest dto.CreateTodoRequest) (dto.TodoResponse, error) {
	todo := models.Todo{
		Title: createTodoRequest.Title,
	}

	u.db.Create(&todo)

	return dto.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}, nil
}
