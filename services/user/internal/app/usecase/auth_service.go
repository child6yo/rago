package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/child6yo/rago/services/auth/internal"
	"github.com/child6yo/rago/services/auth/internal/app/repository"
	"github.com/golang-jwt/jwt/v5"
)

var (
	salt       = "uicn893cion2kds"
	signingKey = "aa1b663d3d5b3baf33523279e65f4cd0"
	tokenTTL   = 3 * time.Hour
)

type tokenClaims struct {
	UserId int  `json:"user_id"`
	Active bool `json:"active"`
	jwt.RegisteredClaims
}

// AuthorizationService имплементирует интерфейс Authorization.
type AuthorizationService struct {
	repo repository.Authorization
}

func (as *AuthorizationService) Register(user internal.User) error {
	user.Password = generatePasswordHash(user.Password)
	err := as.repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (as *AuthorizationService) Login(login, password string) (string, error) {
	user, err := as.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	claims := tokenClaims{
		user.Id,
		user.Active,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signingKey))
}

func (as *AuthorizationService) Auth(accessToken string) error {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return err
	}

	_, ok := token.Claims.(*tokenClaims)
	if !ok {
		return errors.New("token claims are not of type *tokenClaims")
	}

	return nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
