package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// these are the variables

type Config struct {
	PublicHost string
	Port string
	DBUser string
	DBPassword string
	DBAddress string
	DBName string
}

// global variable that be called to initiate the Configuration
var Envs = initConfig()

// return a configuration object
func initConfig() Config {
	godotenv.Load()
	// populate with the variables
	return Config{
		// key
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "ecom"),
	}
}

// a key fallback, just incase the key does not exist, and the value of the key does not exist
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}