package auth

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtSecretKey         = "Jwt_secret_very_secret"
	defaultTokenDuration = 24 * time.Hour
)

type JwtCustomClaim struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// make mathod for jwt token and validation

type JWTService interface {
	GenerateToken(userID uint, email string, name string) (string, error)
	ValidateToken(tokenStr string) (*JwtCustomClaim, error)
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
		UserID:    userId,
		Email:     email,
		Name:      name,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "ticket-booking-app",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *jwtService) ValidateToken(tokenStr string) (*JwtCustomClaim, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaim{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate token failed %w", err)
	}

	if claims, ok := token.Claims.(*JwtCustomClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token Invalid")
}
