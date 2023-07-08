package app

import (
	"fmt"
	"golang-easy-project/internal/config"
	"golang-easy-project/internal/services"
	"golang-easy-project/internal/transport/rest"
)

func Run() {
	fmt.Println("Run initialization whole app")

	config.LoadEnvs()

	db, err := config.DatabaseInit()
	if err != nil {
		panic(err)
	}

	userService := services.GetUserService(db)

	api := rest.GetAPI(userService)
	addr := fmt.Sprintf(":%s", "8080")
	api.Run(addr)
}
