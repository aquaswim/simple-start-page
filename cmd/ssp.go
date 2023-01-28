package main

import (
	"log"
	"simple-start-page/internal"
)

func main() {
	sspServer := internal.NewServer(&internal.Config{
		ListenAddr: ":3000",
		Dev:        true,
	})

	log.Fatalln(sspServer.Start())
}
