package usecases

import "errors"

var (
	ErrTodoNotFound      = errors.New("todo not found")
	ErrTodoAlreadyExists = errors.New("todo already exists")
)
