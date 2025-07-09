package usecases

import "errors"

var (
	ErrProjectNotFound      = errors.New("Projeto não encontrado.")
	ErrProjectAlreadyExists = errors.New("Projeto já existe com esse nome.")
)
