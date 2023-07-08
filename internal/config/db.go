package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func getSqlConfig(config DBConfig) string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", config.Host, config.Port)
	cfg.User = config.User
	cfg.Passwd = config.Password
	cfg.DBName = config.Database
	cfg.Timeout = 5 * time.Second
	cfg.ReadTimeout = 5 * time.Second
	cfg.WriteTimeout = 5 * time.Second
	cfg.ParseTime = true
	return cfg.FormatDSN()
}

func DatabaseInit() (*sql.DB, error) {
	// Retrieve environment variables
	config := GetConfig()
	// Configure MySQL driver with retry settings
	dbConfig := getSqlConfig(config)

	// Retry connecting to the database
	var db *sql.DB
	var err error
	retries := 0
	maxRetries := 10
	for retries < maxRetries {
		db, err = sql.Open("mysql", dbConfig)
		if err == nil {
			break
		}

		retries++
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	// Test the connection
	for retries < maxRetries {
		err = db.Ping()
		if err == nil {
			break
		}

		retries++
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		err := db.Close()
		if err != nil {
			// Log or handle the error appropriately
			log.Println("Error closing the database:", err)
		}
		return nil, err
	}

	fmt.Println("Database connected")

	defer func() {
		err := db.Close()
		if err != nil {
			// Log or handle the error appropriately
			log.Println("Error closing the database:", err)
		}
	}()
	// Return the database connection
	return db, nil
}
