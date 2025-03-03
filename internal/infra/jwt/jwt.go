package jwt

import (
	"lentara-backend/internal/infra/env"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTItf interface {
	GenerateToken(userID uuid.UUID, isAdmin bool) (string, error)
	ValidateToken(token string) (uuid.UUID, bool, error)
}

type JWT struct {
	secretKey   string
	expiredTime int
}

func NewJWT(env *env.Env) *JWT {
	secretKey := env.JWTSecretKey
	expiredTime := env.JWTExpired

	return &JWT{
		secretKey:   secretKey,
		expiredTime: expiredTime,
	}
}

type Claims struct {
	ID      uuid.UUID
	IsAdmin bool
	jwt.RegisteredClaims
}

func (j *JWT) GenerateToken(userID uuid.UUID, isAdmin bool) (string, error) {
	claim := Claims{
		ID:      userID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.expiredTime))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", fiber.NewError(http.StatusInternalServerError, "failed to generate token")
	}

	return tokenString, nil
}

func (j *JWT) ValidateToken(tokenString string) (uuid.UUID, bool, error) {
	var claim Claims

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return uuid.Nil, false, err
	}

	if !token.Valid {
		return uuid.Nil, false, fiber.NewError(http.StatusUnauthorized, "token invalid")
	}

	userID := claim.ID

	return userID, claim.IsAdmin, nil
}
