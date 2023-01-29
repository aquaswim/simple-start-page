package internal

import (
	"fmt"
	"simple-start-page/internal/entities"
	"simple-start-page/internal/pkg"
	"simple-start-page/internal/repositories"
)

func RunDbSeeder(authRepo repositories.Auth, settingRepo repositories.Setting) error {
	fmt.Println("Seed authData: start")
	if err := seedAuth(authRepo); err != nil {
		return err
	}
	fmt.Println("Seed authData: end")

	fmt.Println("Seed setting: start")
	if err := seedSetting(settingRepo); err != nil {
		return err
	}
	fmt.Println("Seed setting: end")
	return nil
}

func seedSetting(repo repositories.Setting) error {
	setting, err := repo.GetSetting()
	if err != nil {
		return err
	}
	if setting != nil && setting.About != "" && setting.SiteTitle != "" {
		return nil
	}
	err = repo.SaveSetting(&entities.Setting{
		SiteTitle:   "My Simple Start Page",
		About:       "Your start page about, you can change this in setting, default username: 'admin' | password: 'password'",
		HideSetting: false,
	})
	if err != nil {
		return err
	}
	return nil
}

func seedAuth(repo repositories.Auth) error {
	data, err := repo.GetAuthData()
	if err != nil {
		return err
	}
	if data != nil && data.Password != "" && data.Username != "" {
		return nil
	}
	password, err := pkg.HashPassword("password")
	if err != nil {
		return err
	}
	return repo.SaveAuthData("admin", password)
}
