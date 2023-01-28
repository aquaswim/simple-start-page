package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"simple-start-page/internal/entities"
	error2 "simple-start-page/internal/error"
	"simple-start-page/internal/pkg"
	"simple-start-page/internal/repositories"
	"time"
)

type auth struct {
	authRepo    repositories.Auth
	tokenSecret string
}

func (a *auth) CekLogin(authData *entities.AuthResult) error {
	token, err := jwt.Parse(authData.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

		}
		return []byte(a.tokenSecret), nil
	})
	if err != nil {
		return &error2.AppError{
			Code:    error2.ErrCodeInvalidCredential,
			Message: err.Error(),
		}
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return &error2.AppError{
			Code:    error2.ErrCodeInvalidCredential,
			Message: "Token has invalid claims",
		}
	}
}

func (a *auth) UpdateAuth(username, password string) error {
	encPasswd, err := pkg.HashPassword(password)
	if err != nil {
		return err
	}
	return a.authRepo.SaveAuthData(username, encPasswd)
}

func (a *auth) Login(username, password string) (*entities.AuthResult, error) {
	authData, err := a.authRepo.GetAuthData()
	if err != nil {
		return nil, err
	}
	if authData.VerifyAuth(username, password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		})
		signedToken, err := token.SignedString([]byte(a.tokenSecret))
		if err != nil {
			return nil, err
		}

		return &entities.AuthResult{
			Token: signedToken,
		}, nil
	}
	return nil, fmt.Errorf("invalid password")
}

type Auth interface {
	Login(username, password string) (*entities.AuthResult, error)
	CekLogin(authData *entities.AuthResult) error
	UpdateAuth(username, password string) error
}

func NewAuthService(secret string, authRepo repositories.Auth) Auth {
	return &auth{
		tokenSecret: secret,
		authRepo:    authRepo,
	}
}
