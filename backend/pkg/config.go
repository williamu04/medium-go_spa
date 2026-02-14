package pkg

import (
	"os"
)

type Config struct {
	RestAPIPort string
	DBDSN       string
	JWTSecret   string
	JWTExpiry   string
	LogLevel    string
	SeedData    string
}

func LoadConfig() *Config {
	return &Config{
		RestAPIPort: os.Getenv("RESTAPI_PORT"),
		DBDSN:       os.Getenv("DB_DSN"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		JWTExpiry:   os.Getenv("JWT_EXPIRY"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
		SeedData:    os.Getenv("SEED_DATA"),
	}
}
