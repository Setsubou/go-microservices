package config

import (
	"log"
	"os"
	"strconv"
)

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("environment key %s is missing", key)
	}

	return os.Getenv(key)
}

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDatabaseUrl() string {
	return getEnvironmentValue("DATABASE_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APP_PORT")

	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %s is not a valid port", portStr)
	}

	return port
}