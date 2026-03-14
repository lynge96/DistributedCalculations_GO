package storage

import (
	"database/sql"
	"errors"
	"log/slog"
	"shared/models"
)

func (s *UserStore) ExistsByUsername(username string) (bool, error) {

	var found string
	query := "SELECT username FROM users WHERE username = ? LIMIT 1"

	err := s.db.QueryRow(query, username).Scan(&found)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		slog.Error("failed to check if user exists", "error", err)
		return false, err
	}

	slog.Info("user exists in database", "username", username)
	return true, nil
}

func (s *UserStore) CreateUser(user models.User) error {

	query := "INSERT INTO users (id, username, password_hash, created_at) VALUES (?, ?, ?, ?)"

	_, err := s.db.Exec(query, user.ID, user.Username, user.PasswordHash, user.CreatedAt)

	slog.Info("user created in database", "username", user.Username)
	return err
}

func (s *UserStore) GetByUsername(username string) (models.User, error) {

	var user models.User
	query := "SELECT id, username, password_hash, created_at FROM users WHERE username = ? LIMIT 1"

	err := s.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Warn("user not found in database", "username", username)
		return user, nil
	}
	if err != nil {
		slog.Error("failed to retrieve user from database", "error", err)
		return user, err
	}

	slog.Info("user retrieved from database", "username", user.Username)
	return user, nil
}
