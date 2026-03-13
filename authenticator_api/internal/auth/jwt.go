package auth

import (
	"authenticator_api/internal/models"
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

type JwtAuth struct {
	secretKey []byte
}

func NewJwtAuth(secret string) *JwtAuth {
	return &JwtAuth{secretKey: []byte(secret)}
}

func (j *JwtAuth) ValidateUser(user models.User, password string) (string, error) {

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := j.generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *JwtAuth) generateToken(user models.User) (string, error) {

	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	slog.Info("generated JWT token for user:", "user_id", user.ID)
	return token.SignedString(j.secretKey)
}

func (j *JwtAuth) ValidateToken(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return j.secretKey, nil
	})
	if err != nil {
		slog.Warn("failed to validate JWT token", "error", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
