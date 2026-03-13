package service

import "authenticator_api/internal/models"

type Storage interface {
	GetByUsername(username string) (models.User, error)
	CreateUser(user models.User) error
	ExistsByUsername(username string) (bool, error)
}

type Authentication interface {
	ValidateUser(user models.User, password string) (string, error)
}

type Service struct {
	storage        Storage
	authentication Authentication
}

func NewService(storage Storage, auth Authentication) *Service {
	return &Service{
		storage:        storage,
		authentication: auth,
	}
}
