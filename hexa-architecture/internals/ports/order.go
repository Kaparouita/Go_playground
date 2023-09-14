package ports

import "github.com/kaparouita/fiber_api/internals/domain"

type OrderRepository interface {
	GetOrder(id int) (*domain.Order, error)
	CreateOrder(order *domain.Order) (*domain.Order, error)
	UpdateOrder(id int, order *domain.Order) (*domain.Order, error)
	DeleteOrder(id int) error
	GetOrders() ([]*domain.Order, error)
}
type OrderService interface {
	GetOrder(id int) (*domain.Order, error)
	CreateOrder(order *domain.Order) (*domain.Order, error)
	UpdateOrder(id int, order *domain.Order) (*domain.Order, error)
	DeleteOrder(id int) error
	GetOrders() ([]*domain.Order, error)
}
