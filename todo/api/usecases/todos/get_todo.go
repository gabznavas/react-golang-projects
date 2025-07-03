package usecases

import (
	dto "api/dto"
	models "api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTodoUsecase struct {
	db *gorm.DB
}

func NewGetTodoUsecase(db *gorm.DB) *GetTodoUsecase {
	return &GetTodoUsecase{db: db}
}

func (u *GetTodoUsecase) GetTodo(ctx *gin.Context, id int) (*dto.TodoResponse, error) {
	var todo models.Todo
	result := u.db.First(&todo, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrTodoNotFound
		}
		return nil, result.Error
	}

	return &dto.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}, nil
}
