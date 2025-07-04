package usecases

import (
	"api/dto"
	"api/models"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type UpdateProjectUsecase struct {
	db                            *gorm.DB
	verifyAttributesUniqueUsecase *VerifyAttributesUniqueUsecase
}

func NewUpdateProjectUsecase(db *gorm.DB, verifyAttributesUniqueUsecase *VerifyAttributesUniqueUsecase) *UpdateProjectUsecase {
	return &UpdateProjectUsecase{
		db:                            db,
		verifyAttributesUniqueUsecase: verifyAttributesUniqueUsecase,
	}
}

func (u *UpdateProjectUsecase) Execute(id uint, params dto.UpdateProjectRequest) error {
	exists, err := u.verifyAttributesUniqueUsecase.CheckIfProjectExistsByName(params.Name, id)
	if err != nil {
		return err
	}
	if exists {
		return ErrProjectAlreadyExists
	}

	project := models.Project{}
	err = u.db.First(&project, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProjectNotFound
		}
		return err
	}
	if params.Name != "" {
		project.Name = strings.TrimSpace(params.Name)
	}
	if params.Description != "" {
		project.Description = strings.TrimSpace(params.Description)
	}
	err = u.db.Save(&project).Error
	if err != nil {
		return err
	}
	return nil
}
