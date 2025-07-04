package usecases

import (
	"api/dto"
	"api/models"
	"errors"

	"gorm.io/gorm"
)

type GetProjectByIdUsecase struct {
	db *gorm.DB
}

func NewGetProjectByIdUsecase(db *gorm.DB) *GetProjectByIdUsecase {
	return &GetProjectByIdUsecase{db: db}
}

func (u *GetProjectByIdUsecase) Execute(id uint) (*dto.ProjectResponse, error) {
	project := models.Project{}
	err := u.db.First(&project, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &dto.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
	}, nil
}
