package usecases

import (
	"api/dto"
	"api/models"

	"gorm.io/gorm"
)

type GetAllProjectsUsecase struct {
	db *gorm.DB
}

func NewGetAllProjectsUsecase(db *gorm.DB) *GetAllProjectsUsecase {
	return &GetAllProjectsUsecase{db: db}
}

func (u *GetAllProjectsUsecase) Execute(search string, offset, limit int) ([]*dto.ProjectResponse, error) {
	var projects []models.Project
	err := u.db.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&projects, "name LIKE ?", "%"+search+"%").
		Error
	if err != nil {
		return nil, err
	}
	projectsResponse := make([]*dto.ProjectResponse, len(projects))
	for i, project := range projects {
		projectsResponse[i] = &dto.ProjectResponse{
			Id:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		}
	}
	return projectsResponse, nil
}
