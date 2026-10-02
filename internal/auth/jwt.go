package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtSecretKey         = "Jwt_secret_very_secret"
	defaultTokenDuration = 24 * time.Hour
)

type JwtCustomClaim struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// make mathod for jwt token and validation

type JWTService interface {
	GenerateToken(userID uint, email string, name string) (string, error)
	// ValidateToken(tokenStr string) (*JwtCustomClaim, error)
}

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTService(secretKey string, tokenDuration time.Duration) JWTService {

	if secretKey == "" {
		secretKey = jwtSecretKey
	}

	return &jwtService{
		secretKey:     secretKey,
		tokenDuration: defaultTokenDuration,
	}
}

// implement generate token and implement methods of interface

func (s *jwtService) GenerateToken(userId uint, email string, name string) (string, error) {
	claims := JwtCustomClaim{
		UserID: userId,
		Name:   name,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ticket-booking-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenStr, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// func (s *jwtService) ValidateToken(tokenStr string) (*JwtCustomClaim, error) {

// }
