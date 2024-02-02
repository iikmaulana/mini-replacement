package main

import (
	"github.com/iikmaulana/mini-replacement/base/config"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Init()

	err = cfg.Server.Start()
	if err != nil {
		panic(err)
	}
}
