package repositories

import (
	"errors"

	"github.com/kaparouita/fiber_api/internals/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func findUser(id int, user *domain.User, db *gorm.DB) error {
	db.Find(&user, "id = ?", id) //gt mporeis na to kaneis auto lol
	//arxizoun apo 1
	if user.Id == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) GetUser(id int) (*domain.User, error) {
	var user domain.User
	if err := findUser(id, &user, r.db); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) DeleteUser(id int) error {
	var user domain.User
	if err := findUser(id, &user, r.db); err != nil {
		return err
	}
	r.db.Delete(&user)
	return nil
}

func (r *UserRepository) UpdateUser(id int, user *domain.User) (*domain.User, error) {
	var finduser domain.User
	if err := findUser(id, &finduser, r.db); err != nil {
		return nil, err
	}
	r.db.Model(&finduser).Updates(&user)
	return &finduser, nil
}
func (r *UserRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	r.db.Find(&users)
	return users, nil
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
