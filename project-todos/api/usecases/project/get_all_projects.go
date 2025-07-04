package usecases

import (
	"api/models"

	"gorm.io/gorm"
)

type GetAllProjectsUsecase struct {
	db *gorm.DB
}

func NewGetAllProjectsUsecase(db *gorm.DB) *GetAllProjectsUsecase {
	return &GetAllProjectsUsecase{db: db}
}

func (u *GetAllProjectsUsecase) Execute() ([]models.Project, error) {
	var projects []models.Project
	err := u.db.Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}
