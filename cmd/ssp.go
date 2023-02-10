package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"simple-start-page/internal"
	"simple-start-page/internal/docker"
	"simple-start-page/internal/handlers"
	"simple-start-page/internal/repositories"
	"simple-start-page/internal/services"
	"simple-start-page/pkg"
)

var dbPath string
var authSecret string
var listenAddr string
var isDev bool
var skipSeeder bool
var enableDockerIntegration bool

func init() {
	flag.StringVar(&dbPath, "db", "", "DB File Path")
	flag.StringVar(&authSecret, "secret", "", "JWT secret")
	flag.StringVar(&listenAddr, "listen", "127.0.0.1:3000", "Listen address")
	flag.BoolVar(&isDev, "dev", false, "Dev Mode")
	flag.BoolVar(&skipSeeder, "skipSeed", false, "Skip seeding")
	flag.BoolVar(&enableDockerIntegration, "docker", false, "Enable docker integration")
}

func main() {
	flag.Parse()
	if authSecret == "" {
		fmt.Println("Error: Need to set JWT secret")
		flag.Usage()
		os.Exit(-1)
	}
	if dbPath == "" {
		fmt.Println("Error: Need to set DB Path")
		flag.Usage()
		os.Exit(-1)
	}

	db, err := pkg.NewDB(dbPath)
	if err != nil {
		log.Fatalln("Error opening db", err)
	}
	defer db.Close()

	authRepo := repositories.NewAuthRepo(db)
	settingRepo := repositories.NewSettingRepo(db)
	if !skipSeeder {
		err := internal.RunDbSeeder(authRepo, settingRepo)
		if err != nil {
			log.Fatalln("Error running seeder", err)
		}
	}
	dockerClient, err := (func() (docker.Docker, error) {
		if enableDockerIntegration {
			return docker.NewDockerIntegration()
		} else {
			return docker.NewDockerDummyClient(), nil
		}
	})()
	if err != nil {
		log.Fatalln("Error creating docker client: ", err)
	}

	authSvc := services.NewAuthService(authSecret, authRepo)
	settingSvc := services.NewSettingService(settingRepo, dockerClient)
	apiHandler := handlers.NewApiHandler(authSvc, settingSvc)
	sspServer := internal.NewServer(&internal.Config{
		ListenAddr: ":3000",
		Dev:        isDev,
	}, apiHandler)

	log.Fatalln(sspServer.Start())
}
