package usecases

import (
	"api/dto"
	"api/models"

	"gorm.io/gorm"
)

type CreateProjectUsecase struct {
	db *gorm.DB
}

func NewCreateProjectUsecase(db *gorm.DB) *CreateProjectUsecase {
	return &CreateProjectUsecase{db: db}
}

func (u *CreateProjectUsecase) Execute(params *dto.CreateProjectRequest) (*dto.ProjectResponse, error) {
	project := models.Project{
		Name:        params.Name,
		Description: params.Description,
	}
	err := u.db.Create(&project).Error
	if err != nil {
		return nil, err
	}
	return &dto.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
	}, nil
}
