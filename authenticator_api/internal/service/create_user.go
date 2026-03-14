package service

import (
	"errors"
	"log/slog"
	"shared/models"
)

func (s *Service) CreateUser(username, password string) error {

	exists, err := s.storage.ExistsByUsername(username)
	if err != nil {
		slog.Error("failed to check if user exists", "error", err)
		return err
	}

	if exists {
		slog.Warn("user already exists", "username", username)
		return errors.New("user already exists")
	}

	user, err := models.NewUser(username, password)
	if err != nil {
		slog.Error("failed to create user model", "error", err)
		return err
	}

	err = s.storage.CreateUser(user)
	if err != nil {
		slog.Error("failed to store user", "error", err)
		return err
	}

	return nil
}
