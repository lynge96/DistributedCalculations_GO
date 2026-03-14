package auth

import (
	auth "shared/auth"
	"shared/models"

	"golang.org/x/crypto/bcrypt"
)

type Authenticator struct {
	jwt *auth.JwtAuth
}

func NewAuthenticator(jwt *auth.JwtAuth) *Authenticator {
	return &Authenticator{jwt: jwt}
}

func (a *Authenticator) ValidateUser(user models.User, password string) (string, error) {

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := a.jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
