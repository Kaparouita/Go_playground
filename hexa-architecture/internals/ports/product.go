package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaparouita/fiber_api/internals/domain"
)

type ProductRepository interface {
	GetProduct(id int) (*domain.Product, error)
	CreateProduct(product *domain.Product) error
	UpdateProduct(id int, product *domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
	GetProducts() ([]*domain.Product, error)
}

type ProductService interface {
	GetProduct(id int) (*domain.Product, error)
	CreateProduct(product *domain.Product) error
	UpdateProduct(id int, product *domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
	GetProducts() ([]*domain.Product, error)
}

type ProductHandlers interface {
	GetProduct(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
}
