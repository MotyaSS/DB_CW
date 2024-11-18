package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HttpServer `yaml:"http-server" env-required:"true"`
	Database   `yaml:"database"`
}

type HttpServer struct {
	Address string        `yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type Database struct {
	Address  string `yaml:"address" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	configPath, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		return nil, fmt.Errorf("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file doesn't exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("error occured while reading config: %s", err.Error())
	}
	return &cfg, nil
}
