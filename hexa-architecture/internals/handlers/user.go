package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaparouita/fiber_api/internals/domain"
	"github.com/kaparouita/fiber_api/internals/ports"
)

type UserHandlers struct {
	UserService ports.UserService
}

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{UserService: userService}
}

func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := h.UserService.CreateUser(&user); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).JSON(user)
}

func (h *UserHandlers) GetUser(c *fiber.Ctx) error {
	var user *domain.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if user, err = h.UserService.GetUser(id); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(&user)
}

func (h *UserHandlers) UpdateUser(c *fiber.Ctx) error {
	var user domain.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var responseUser *domain.User
	responseUser, err = h.UserService.UpdateUser(id, &user)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).JSON(&responseUser)
}

func (h *UserHandlers) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := h.UserService.DeleteUser(id); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully deleted")
}

func (h *UserHandlers) GetUsers(c *fiber.Ctx) error {
	var users []*domain.User
	var err error
	users, err = h.UserService.GetUsers()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(200).JSON(&users)
}
