package usecases

import (
	"api/models"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type VerifyAttributesUniqueUsecase struct {
	db *gorm.DB
}

func NewVerifyAttributesUniqueUsecase(db *gorm.DB) *VerifyAttributesUniqueUsecase {
	return &VerifyAttributesUniqueUsecase{db: db}
}

func (u *VerifyAttributesUniqueUsecase) CheckIfProjectExistsByName(name string, id uint) (bool, error) {
	alreadyExists := models.Project{}
	err := u.db.Where("LOWER(name) = ?", strings.ToLower(strings.TrimSpace(name))).First(&alreadyExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return alreadyExists.ID != 0 && alreadyExists.ID != id, nil
}
