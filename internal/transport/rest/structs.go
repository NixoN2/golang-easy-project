package rest

import (
	"github.com/gin-gonic/gin"
	"golang-easy-project/internal/services"
)

type API struct {
	router      *gin.Engine
	userService *services.UserService
}
