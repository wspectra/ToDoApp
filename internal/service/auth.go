package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wspectra/api_server/internal/pkg/utils"
	"github.com/wspectra/api_server/internal/repository"
	"github.com/wspectra/api_server/internal/structure"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

const (
	signedKey = "eodjfejngkdjfgnkdefajk"
	tokenTTl  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) AddNewUser(user structure.User) error {
	user.Password = utils.GeneratePasswordHash(user.Password)
	return a.repo.AddNewUser(user)
}

func (a *AuthService) GetToken(user structure.SignInUser) (string, error) {
	user.Password = utils.GeneratePasswordHash(user.Password)
	userId, err := a.repo.AuthorizeUser(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	return token.SignedString([]byte(signedKey))
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signedKey), nil
	})
	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("token claims are not of type")
	}

	return claims.UserId, nil
}
