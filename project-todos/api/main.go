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
	verifyAttributesUniqueUsecase := projectUsecase.NewVerifyAttributesUniqueUsecase(db)
	getAllProjectsUsecase := projectUsecase.NewGetAllProjectsUsecase(db)
	createProjectUsecase := projectUsecase.NewCreateProjectUsecase(db, verifyAttributesUniqueUsecase)
	getProjectByIdUsecase := projectUsecase.NewGetProjectByIdUsecase(db)
	updateProjectUsecase := projectUsecase.NewUpdateProjectUsecase(db, verifyAttributesUniqueUsecase)
	deleteProjectUsecase := projectUsecase.NewDeleteProjectUsecase(db)

	// Controllers
	getAllProjectsController := projectController.NewGetAllProjectsController(getAllProjectsUsecase)
	createProjectController := projectController.NewCreateProjectController(createProjectUsecase)
	getProjectByIdController := projectController.NewGetProjectByIdController(getProjectByIdUsecase)
	updateProjectController := projectController.NewUpdateProjectController(updateProjectUsecase)
	deleteProjectController := projectController.NewDeleteProjectController(deleteProjectUsecase)

	// Routes
	router := gin.Default()
	router.GET("api/v1/project", getAllProjectsController.GetAllProjects)
	router.POST("api/v1/project", createProjectController.CreateProject)
	router.GET("api/v1/project/:id", getProjectByIdController.GetProjectById)
	router.PUT("api/v1/project/:id", updateProjectController.UpdateProject)
	router.DELETE("api/v1/project/:id", deleteProjectController.DeleteProject)
	router.Run(":8080")
}
