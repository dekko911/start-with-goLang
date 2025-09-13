package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost, Port, DBUser, DBPassword, DBAddress, DBName string
}

var Env = initConfig()

// Getting config at env file or somewhere else.
func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")

	return Config{
		PublicHost: os.Getenv("PUBLIC_HOST"),
		Port:       os.Getenv("PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddress:  fmt.Sprintf("%s:%s", DBHost, DBPort),
		DBName:     os.Getenv("DB_NAME"),
	}
}

// This method use for getting config env manually.
// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}

// 	return fallback
// }
