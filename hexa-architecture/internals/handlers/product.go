package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaparouita/fiber_api/internals/domain"
	"github.com/kaparouita/fiber_api/internals/ports"
)

type ProductHandler struct {
	ProductService ports.ProductService
}

func NewProductHandler(productService ports.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := h.ProductService.CreateProduct(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var product domain.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var responseProduct *domain.Product
	responseProduct, err = h.ProductService.UpdateProduct(id, &product)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(responseProduct)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := h.ProductService.DeleteProduct(id); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully deleted product")
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product *domain.Product
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	product, err = h.ProductService.GetProduct(id)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(&product)
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.ProductService.GetProducts()
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(&products)
}
