package rest

func (api *API) RegisterUserRoutes() {
	userGroup := api.router.Group("/users")
	userGroup.GET("/", api.userService.GetUserList)
	userGroup.POST("/", api.userService.CreateUser)
}
