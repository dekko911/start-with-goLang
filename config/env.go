package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost, Port, DBUser, DBPassword, DBAddress, DBName, JWTSecret string
	JWTExpirationInSeconds                                             int64
}

var Env = initConfig()

// Getting config at env file or somewhere else.
func initConfig() Config {
	godotenv.Load()

	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")

	return Config{
		PublicHost:             os.Getenv("PUBLIC_HOST"),
		Port:                   os.Getenv("PORT"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBAddress:              fmt.Sprintf("%s:%s", DBHost, DBPort),
		DBName:                 os.Getenv("DB_NAME"),
		JWTExpirationInSeconds: getEnvAsInt(os.Getenv("JWT_EXP"), 3600*24*7),
		JWTSecret:              os.Getenv("JWT_SECRET"),
	}
}

// This method use for getting config env manually.
// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}

// 	return fallback
// }

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}
	return fallback
}
