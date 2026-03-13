package storage

import (
	"authenticator_api/internal/models"
	"database/sql"
	"errors"
)

func (s *UserStore) ExistsByUsername(username string) (bool, error) {

	var found string
	query := "SELECT username FROM users WHERE username = ? LIMIT 1"

	err := s.db.QueryRow(query, username).Scan(&found)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *UserStore) CreateUser(user models.User) error {

	query := "INSERT INTO users (id, username, password_hash, created_at) VALUES (?, ?, ?, ?)"

	_, err := s.db.Exec(query, user.ID, user.Username, user.PasswordHash, user.CreatedAt)
	return err
}

func (s *UserStore) GetByUsername(username string) (models.User, error) {

	var user models.User
	query := "SELECT id, username, password_hash, created_at FROM users WHERE username = ? LIMIT 1"

	err := s.db.QueryRow(query, username).
		Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return user, nil
	}
	if err != nil {
		return user, err
	}
	return user, nil
}
