package dto

import "time"

type CreateTodoRequest struct {
	Title string `json:"title" binding:"required" message:"Title is required"`
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateTodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
