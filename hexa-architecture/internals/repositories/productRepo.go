package repositories

import (
	"errors"

	"github.com/kaparouita/fiber_api/internals/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func findProduct(id int, product *domain.Product, db *gorm.DB) error {
	db.Find(&product, "id = ?", id)
	if product.Id == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (r *productRepository) CreateProduct(product *domain.Product) error {
	if err := r.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) GetProduct(id int) (*domain.Product, error) {
	var product domain.Product
	if err := findProduct(id, &product, repo.db); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) UpdateProduct(id int, product *domain.Product) (*domain.Product, error) {
	var findproduct domain.Product
	if err := findProduct(id, &findproduct, r.db); err != nil {
		return nil, err
	}
	r.db.Model(&findproduct).Updates(&product)
	return &findproduct, nil
}

func (r *productRepository) DeleteProduct(id int) error {
	var product domain.Product
	if err := findProduct(id, &product, r.db); err != nil {
		return err
	}
	r.db.Delete(&product)
	return nil
}

func (r *productRepository) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	r.db.Find(&products)
	return products, nil
}
