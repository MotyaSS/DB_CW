package main

import (
	"log"

	cfg "github.com/MotyaSS/DB_CW/pkg/config"
	hnd "github.com/MotyaSS/DB_CW/pkg/handler"
	srvr "github.com/MotyaSS/DB_CW/pkg/server"
	srvc "github.com/MotyaSS/DB_CW/pkg/service"
	strg "github.com/MotyaSS/DB_CW/pkg/storage"
	"github.com/joho/godotenv"
)

func main() {
	config := setupConfig()
	log.Println(config)
	storage := strg.New()
	service := srvc.New(storage)
	handler := hnd.New(service)
	server := srvr.New(config.HttpServer.Address, handler.InitRouter())
	server.Run()
}

func setupConfig() *cfg.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config, err := cfg.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}
