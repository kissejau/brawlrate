package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey string
	Host   string
	Port   string
}

func NewConfig() (*Config, error) {
	config := &Config{}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	envPath := filepath.Join(wd, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}

	apiKey := os.Getenv("API_KEY")
	if len(apiKey) == 0 {
		return nil, fmt.Errorf("API_KEY field not found in .env")
	}

	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if host == "" || port == "" {
		return nil, fmt.Errorf("host and port must be added to .env")
	}

	config.Host = host
	config.Port = port
	config.ApiKey = apiKey
	return config, nil
}
