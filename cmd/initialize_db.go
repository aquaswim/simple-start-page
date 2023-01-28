package main

import (
	"flag"
	"fmt"
	"log"
	"simple-start-page/internal/entities"
	pkgInternal "simple-start-page/internal/pkg"
	"simple-start-page/internal/repositories"
	"simple-start-page/pkg"
)

var dbFile string
var isHelp bool

func init() {
	flag.StringVar(&dbFile, "db", "./test.db", "dbfile path")
	flag.BoolVar(&isHelp, "help", false, "Show help")
}

func main() {
	flag.Parse()
	if isHelp {
		flag.Usage()
		return
	}

	fmt.Println("Opening db", dbFile)
	db, err := pkg.NewDB(dbFile)
	if err != nil {
		log.Fatalln("Error opening db", err)
	}
	defer db.Close()
	// init repo
	authRepo := repositories.NewAuthRepo(db)
	settingRepo := repositories.NewSettingRepo(db)

	// populate data
	password, _ := pkgInternal.HashPassword("password")
	authRepo.SaveAuthData("admin", password)
	settingRepo.SaveSetting(&entities.Setting{
		SiteTitle:   "My Simple Site",
		About:       "Just simple site for doing anything",
		HideSetting: true,
	})
	settingRepo.SaveListUrls(&[]entities.Link{
		{Name: "Search Engine", Url: "https://duckduckgo.com", Icon: "search"},
	})
	fmt.Println("db reinitialized")
}
