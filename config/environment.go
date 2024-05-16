package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	DbServer   string
	DbUser     string
	DbPassword string
	DbPort     int
	DbDatabase string
}

var Env = newEnvironment()

func newEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Environment{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbDatabase: os.Getenv("DB_DATABASE"),
		DbServer:   os.Getenv("DB_SERVER"),
		DbPort:     getEnvironmentVariableAsInt("DB_PORT"),
	}
}

func getEnvironmentVariableAsInt(key string) int {

	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0
	}
	return value
}
