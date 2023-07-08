package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Print environment variables
	fmt.Println("MYSQL_ROOT_PASSWORD:", os.Getenv("MYSQL_ROOT_PASSWORD"))
	fmt.Println("MYSQL_DATABASE:", os.Getenv("MYSQL_DATABASE"))
	fmt.Println("MYSQL_DB_USER:", os.Getenv("MYSQL_DB_USER"))
	fmt.Println("MYSQL_DB_HOST:", os.Getenv("MYSQL_DB_HOST"))
	fmt.Println("MYSQL_DB_PORT:", os.Getenv("MYSQL_DB_PORT"))
}

func GetConfig() DBConfig {
	config := DBConfig{
		User:     os.Getenv("MYSQL_DB_USER"),
		Password: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Host:     os.Getenv("MYSQL_DB_HOST"),
		Port:     os.Getenv("MYSQL_DB_PORT"),
	}
	return config
}
