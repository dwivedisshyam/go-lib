package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
	GetDefault(key, value string) string
}

func New() Config {
	cp := new(configProvider)

	read("./configs/.env")

	env := os.Getenv("APP_ENV")
	if env == "" {
		return cp
	}

	read("./configs/." + env + ".env")

	return cp
}

func read(file string) {
	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("Error loading .env file. [%v]", err)
	}
}

func NewMockConfig() Config {
	return &MockConfig{}
}
