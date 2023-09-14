package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kaparouita/fiber_api/internals/domain"
)

type UserService interface {
	GetUser(id int) (*domain.User, error)
	CreateUser(user *domain.User) error
	UpdateUser(id int, user *domain.User) (*domain.User, error)
	DeleteUser(id int) error
	GetUsers() ([]*domain.User, error)
}

type UserRepository interface {
	GetUser(id int) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(id int, user *domain.User) (*domain.User, error)
	DeleteUser(id int) error
	GetUsers() ([]*domain.User, error)
}

type UserHandlers interface {
	GetUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
}
