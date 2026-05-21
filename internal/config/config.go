package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	PostgresURL string
	S3Endpoint  string
	S3AccessKey string
	S3SecretKey string
	S3Bucket    string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	return Config{
		Port:        os.Getenv("PORT"),
		PostgresURL: os.Getenv("POSTGRES_URL"),

		S3Endpoint:  os.Getenv("S3_ENDPOINT"),
		S3AccessKey: os.Getenv("S3_ACCESS_KEY"),
		S3SecretKey: os.Getenv("S3_SECRET_KEY"),
		S3Bucket:    os.Getenv("S3_BUCKET"),
	}
}
