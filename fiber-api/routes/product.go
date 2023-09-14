package routes

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kaparouita/fiber_api/database"
	"github.com/kaparouita/fiber_api/models"
)

type Product struct {
	Id           uint `json:"id"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

// TRANSLATE PRODUCT
func CreateResponseProduct(productModel models.Product) Product {
	return Product{Id: productModel.Id, CreatedAt: productModel.CreatedAt, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

// CREATE PRODUCT
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

// GET PRODUCTS
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProducts = append(responseProducts, CreateResponseProduct(product))
	}
	return c.Status(200).JSON(responseProducts)
}

// GET PRODUCT
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findProductById(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

// Given an ID check if product exists
func findProductById(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.Id == 0 {
		return errors.New("Product not found")
	}
	return nil
}

// UPDATE PRODUCT
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//GET THE PRODUCT FROM DB IF EXISTS
	if err := findProductById(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	//GET NEW DATA
	type UpdateProductResponse struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}
	var response UpdateProductResponse
	if err := c.BodyParser(&response); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//UPDATE PRODUCT
	product.Name = response.Name
	product.SerialNumber = response.SerialNumber

	database.Database.Db.Save(&product)

	productResponse := CreateResponseProduct(product)
	return c.Status(200).JSON(productResponse)
}

// DELETE PRODUCT
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//GET THE PRODUCT FROM DB IF EXISTS
	if err := findProductById(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	database.Database.Db.Delete(&product)

	return c.Status(200).SendString("Product deleted successfully")
}
