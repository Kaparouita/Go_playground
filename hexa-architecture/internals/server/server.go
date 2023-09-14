package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kaparouita/fiber_api/internals/ports"
)

type Server struct {
	UserHandlers    ports.UserHandlers
	ProductHandlers ports.ProductHandlers
}

func NewServer(userHandlers ports.UserHandlers,
	productHandlers ports.ProductHandlers) *Server {
	return &Server{
		UserHandlers:    userHandlers,
		ProductHandlers: productHandlers,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	api := app.Group("/api")

	//UserRoutes
	userRoutes := api.Group("/users")
	userRoutes.Post("/", s.UserHandlers.CreateUser)
	userRoutes.Get("/:id", s.UserHandlers.GetUser)
	userRoutes.Get("/", s.UserHandlers.GetUsers)
	userRoutes.Put("/:id", s.UserHandlers.UpdateUser)
	userRoutes.Delete("/:id", s.UserHandlers.DeleteUser)

	//ProductRoutes
	productRoutes := api.Group("/products")
	productRoutes.Post("/", s.ProductHandlers.CreateProduct)
	productRoutes.Get("/:id", s.ProductHandlers.GetProduct)
	productRoutes.Get("/", s.ProductHandlers.GetProducts)
	productRoutes.Put("/:id", s.ProductHandlers.UpdateProduct)
	productRoutes.Delete("/:id", s.ProductHandlers.DeleteProduct)

	log.Fatal(app.Listen(":3000"))

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
