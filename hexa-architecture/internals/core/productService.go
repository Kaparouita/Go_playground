package core

import (
	"github.com/kaparouita/fiber_api/internals/domain"
	"github.com/kaparouita/fiber_api/internals/ports"
)

type ProductService struct {
	productRepository ports.ProductRepository
}

func NewProductService(productRepository ports.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) GetProduct(id int) (*domain.Product, error) {
	return s.productRepository.GetProduct(id)
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	return s.productRepository.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(id int, product *domain.Product) (*domain.Product, error) {
	return s.productRepository.UpdateProduct(id, product)
}
func (s *ProductService) DeleteProduct(id int) error {
	return s.productRepository.DeleteProduct(id)
}

func (s *ProductService) GetProducts() ([]*domain.Product, error) {
	return s.productRepository.GetProducts()
}
