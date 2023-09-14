package main

import (
	"github.com/kaparouita/fiber_api/internals/core"
	"github.com/kaparouita/fiber_api/internals/handlers"
	"github.com/kaparouita/fiber_api/internals/repositories"
	"github.com/kaparouita/fiber_api/internals/server"
)

func main() {
	db := repositories.ConnectDb()
	//REPOS
	userRepo := repositories.NewUserRepository(db.Db)
	productRepo := repositories.NewProductRepository(db.Db)
	//SERVICES
	userService := core.NewUserService(userRepo)
	productService := core.NewProductService(productRepo)
	//HANDLERS
	UserHandlers := handlers.NewUserHandlers(userService)
	ProductHandlers := handlers.NewProductHandler(productService)
	//SERVER
	httpServer := server.NewServer(UserHandlers, ProductHandlers)

	httpServer.Initialize()
}
