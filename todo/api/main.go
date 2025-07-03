package main

import (
	todoController "api/controllers/todos"
	models "api/models"
	todoUsecases "api/usecases/todos"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// abre o banco de dados postgres
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5440/postgres"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// cria as tabelas no banco de dados
	db.AutoMigrate(&models.Todo{})

	// cria os usecases
	createTodoUsecase := todoUsecases.NewCreateTodoUsecase(db)
	getTodoUsecase := todoUsecases.NewGetTodoUsecase(db)
	deleteTodoUsecase := todoUsecases.NewDeleteTodoUsecase(db)
	updateTodoUsecase := todoUsecases.NewUpdateTodoUsecase(db)
	getAllTodosUsecase := todoUsecases.NewGetAllTodosUsecase(db)

	// cria as rotas do gin e associa as rotas aos usecases
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/api/v1/todos", todoController.NewCreateTodoController(createTodoUsecase).CreateTodo)
	router.GET("/api/v1/todos/:id", todoController.NewGetTodoController(getTodoUsecase).GetTodo)
	router.DELETE("/api/v1/todos/:id", todoController.NewDeleteTodoController(deleteTodoUsecase).DeleteTodo)
	router.PUT("/api/v1/todos/:id", todoController.NewUpdateTodoController(updateTodoUsecase).UpdateTodo)
	router.GET("/api/v1/todos", todoController.NewGetAllTodosController(getAllTodosUsecase).GetAllTodos)

	// inicia o servidor
	router.Run(":8080")
}
