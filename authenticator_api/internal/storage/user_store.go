package storage

import "authenticator_api/internal/models"

func ExistsByUsername(username string) (bool, error) {
	return false, nil
}

func CreateUser(user models.User) error {
	return nil
}

func GetByUsername(username string) (models.User, error) {
	return models.User{}, nil
}
