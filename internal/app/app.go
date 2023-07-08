package app

import (
	"fmt"
	"golang-easy-project/internal/config"
	"log"
)

func Run() {
	fmt.Println("Run initialization whole app")

	config.LoadEnvs()

	db, err := config.DatabaseInit()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			// Log or handle the error appropriately
			log.Println("Error closing the database:", err)
		}
	}()
}
