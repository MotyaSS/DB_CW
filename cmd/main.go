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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := setupConfig()
	db, err := strg.NewPostgresDB(strg.Config{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		DBName:   config.Database.DBName,
		SSLMode:  config.Database.SSLMode,
	})
	storage := strg.New(db)
	service := srvc.New(storage)
	handler := hnd.New(service)
	server := srvr.New(&config.HttpServer, handler.InitRouter())
	server.Run()
}

func setupConfig() *cfg.Config {
	config, err := cfg.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}
