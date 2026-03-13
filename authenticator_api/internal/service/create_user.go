package service

import (
	"authenticator_api/internal/models"
	"errors"
)

func (s *Service) CreateUser(username, password string) error {

	exists, err := s.storage.ExistsByUsername(username)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	user, err := models.NewUser(username, password)
	if err != nil {
		return err
	}
	
	err = s.storage.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
