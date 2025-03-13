package main

import (
	"log/slog"
	"os"

	cfg "github.com/MotyaSS/DB_CW/pkg/config"
	hnd "github.com/MotyaSS/DB_CW/pkg/handler"
	srvr "github.com/MotyaSS/DB_CW/pkg/server"
	srvc "github.com/MotyaSS/DB_CW/pkg/service"
	strg "github.com/MotyaSS/DB_CW/pkg/storage"
	"github.com/joho/godotenv"
)

// 071ge87tadv123
func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	config := setupConfig()

	db, err := strg.NewPostgresDB(
		strg.Config{
			Host:     config.Database.Host,
			Port:     config.Database.Port,
			Username: config.Database.Username,
			Password: config.Database.Password,
			DBName:   config.Database.DBName,
			SSLMode:  config.Database.SSLMode,
		})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	storage := strg.New(db)

	service := srvc.New(storage, config.Database)
	handler := hnd.New(service)
	server := srvr.New(
		config.Address,
		config.Timeout,
		handler.InitRouter(),
	)

	server.Run()
}

func setupConfig() *cfg.Config {
	config, err := cfg.LoadConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return config
}
