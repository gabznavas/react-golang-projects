package main

import (
	projectController "api/controllers/project"
	"api/models"
	projectUsecase "api/usecases/project"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to database
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5440/postgres"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate models
	db.AutoMigrate(&models.Project{})

	// Usecases
	getAllProjectsUsecase := projectUsecase.NewGetAllProjectsUsecase(db)
	createProjectUsecase := projectUsecase.NewCreateProjectUsecase(db)
	getProjectByIdUsecase := projectUsecase.NewGetProjectByIdUsecase(db)

	// Controllers
	getAllProjectsController := projectController.NewGetAllProjectsController(getAllProjectsUsecase)
	createProjectController := projectController.NewCreateProjectController(createProjectUsecase)
	getProjectByIdController := projectController.NewGetProjectByIdController(getProjectByIdUsecase)

	// Routes
	router := gin.Default()
	router.GET("api/v1/project", getAllProjectsController.GetAllProjects)
	router.POST("api/v1/project", createProjectController.CreateProject)
	router.GET("api/v1/project/:id", getProjectByIdController.GetProjectById)
	router.Run(":8080")
}
