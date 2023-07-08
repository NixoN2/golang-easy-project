package rest

import (
	"github.com/gin-gonic/gin"
	"golang-easy-project/internal/services"
	"log"
)

func (api *API) registerRoutes() {
	api.RegisterUserRoutes()
}

func GetAPI(userService *services.UserService) *API {
	// Create a new Gin router
	router := gin.Default()

	// Initialize the API with the router and user services
	api := &API{
		router:      router,
		userService: userService,
	}

	// Register routes
	api.registerRoutes()

	return api
}

func (api *API) Run(addr string) {
	// Start the server
	err := api.router.Run(addr)
	if err != nil {
		log.Fatal("Failed to run the server:", err)
	}
}
