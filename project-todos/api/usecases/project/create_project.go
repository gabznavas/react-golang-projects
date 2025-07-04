package usecases

import (
	"api/dto"
	"api/models"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CreateProjectUsecase struct {
	db                            *gorm.DB
	verifyAttributesUniqueUsecase *VerifyAttributesUniqueUsecase
}

func NewCreateProjectUsecase(db *gorm.DB, verifyAttributesUniqueUsecase *VerifyAttributesUniqueUsecase) *CreateProjectUsecase {
	return &CreateProjectUsecase{
		db:                            db,
		verifyAttributesUniqueUsecase: verifyAttributesUniqueUsecase,
	}
}

func (u *CreateProjectUsecase) Execute(params *dto.CreateProjectRequest) (*dto.ProjectResponse, error) {
	exists, err := u.verifyAttributesUniqueUsecase.CheckIfProjectExistsByName(params.Name, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("project already exists")
	}

	project := models.Project{
		Name:        strings.TrimSpace(params.Name),
		Description: strings.TrimSpace(params.Description),
	}
	err = u.db.Create(&project).Error
	if err != nil {
		return nil, err
	}
	return &dto.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
	}, nil
}
