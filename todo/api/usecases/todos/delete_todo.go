package usecases

import (
	models "api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteTodoUsecase struct {
	db *gorm.DB
}

func NewDeleteTodoUsecase(db *gorm.DB) *DeleteTodoUsecase {
	return &DeleteTodoUsecase{db: db}
}

func (u *DeleteTodoUsecase) DeleteTodo(ctx *gin.Context, id int) error {
	var todo models.Todo
	result := u.db.First(&todo, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ErrTodoNotFound
		}
		return result.Error
	}

	u.db.Delete(&todo)

	return nil
}
