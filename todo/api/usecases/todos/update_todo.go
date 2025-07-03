package usecases

import (
	dto "api/dto"
	models "api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateTodoUsecase struct {
	db *gorm.DB
}

func NewUpdateTodoUsecase(db *gorm.DB) *UpdateTodoUsecase {
	return &UpdateTodoUsecase{db: db}
}

func (u *UpdateTodoUsecase) UpdateTodo(ctx *gin.Context, id int, updateTodoRequest dto.UpdateTodoRequest) error {
	var todo models.Todo
	result := u.db.First(&todo, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ErrTodoNotFound
		}
		return result.Error
	}

	var existingTodo models.Todo
	result = u.db.Where("title = $1", updateTodoRequest.Title).First(&existingTodo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}
	if existingTodo.ID != 0 && existingTodo.ID != todo.ID {
		return ErrTodoAlreadyExists
	}

	todo.Title = updateTodoRequest.Title
	todo.Completed = updateTodoRequest.Completed
	u.db.Save(&todo)

	return nil
}
