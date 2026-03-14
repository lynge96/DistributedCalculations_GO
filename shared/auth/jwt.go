package auth

import (
	"errors"
	"log/slog"
	"shared/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAuth struct {
	secretKey []byte
}

func NewJwtAuth(secret string) *JwtAuth {
	return &JwtAuth{secretKey: []byte(secret)}
}

func (j *JwtAuth) GenerateToken(user models.User) (string, error) {

	claims := models.Claims{
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

func (j *JwtAuth) ValidateToken(tokenString string) (*models.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (any, error) {
		return j.secretKey, nil
	})
	if err != nil {
		slog.Warn("failed to validate JWT token", "error", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
