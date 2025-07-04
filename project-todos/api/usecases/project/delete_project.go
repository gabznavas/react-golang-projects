package usecases

import (
	"api/models"
	"errors"

	"gorm.io/gorm"
)

type DeleteProjectUsecase struct {
	db *gorm.DB
}

func NewDeleteProjectUsecase(db *gorm.DB) *DeleteProjectUsecase {
	return &DeleteProjectUsecase{db: db}
}

func (u *DeleteProjectUsecase) Execute(id uint) error {
	project := models.Project{}
	err := u.db.First(&project, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProjectNotFound
		}
		return err
	}
	return u.db.Delete(&project).Error
}
