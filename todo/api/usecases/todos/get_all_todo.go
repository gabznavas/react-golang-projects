package usecases

import (
	dto "api/dto"
	models "api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetAllTodosUsecase struct {
	db *gorm.DB
}

func NewGetAllTodosUsecase(db *gorm.DB) *GetAllTodosUsecase {
	return &GetAllTodosUsecase{db: db}
}

func (u *GetAllTodosUsecase) GetAllTodos(ctx *gin.Context, offset int, limit int) ([]dto.TodoResponse, error) {
	var todos []models.Todo
	u.db.Offset(offset).Limit(limit).Find(&todos)

	var todoResponses []dto.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, dto.TodoResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		})
	}

	return todoResponses, nil
}
