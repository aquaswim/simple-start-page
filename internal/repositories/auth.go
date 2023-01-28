package repositories

import (
	"simple-start-page/internal/entities"
	"simple-start-page/pkg"
)

const (
	keyAuth = "auth"
)

type auth struct {
	db pkg.Database
}

func (a *auth) SaveAuthData(username, encryptedPassword string) error {
	au := entities.Auth{
		Username: username,
		Password: encryptedPassword,
	}
	return a.db.SaveObject(keyAuth, au)

}

func (a *auth) GetAuthData() (*entities.Auth, error) {
	au := new(entities.Auth)
	err := a.db.GetObject(keyAuth, au)
	if err != nil {
		return nil, err
	}
	return au, nil
}

type Auth interface {
	GetAuthData() (*entities.Auth, error)
	SaveAuthData(username, encryptedPassword string) error
}

func NewAuthRepo(db pkg.Database) Auth {
	return &auth{
		db: db,
	}
}
