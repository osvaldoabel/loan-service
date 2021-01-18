package services

import (
	"osvaldoabel/smartmei-service/src/entities"
)

//Writer user writer
type Writer interface {
	Create(e *entities.User) (*entities.User, error)
}

//Repository interface
type Repository interface {
	Writer
}
