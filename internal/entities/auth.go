package entities

import "simple-start-page/internal/pkg"

type Auth struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResult struct {
	Token string `json:"token" validate:"required,jwt"`
}

func (a *Auth) VerifyAuth(username, password string) bool {
	if username != a.Username {
		return false
	}
	err := pkg.ComparePassword(a.Password, password)
	return err == nil
}
