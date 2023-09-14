package core

import (
	"github.com/kaparouita/fiber_api/internals/domain"
	"github.com/kaparouita/fiber_api/internals/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUser(id int) (*domain.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) CreateUser(user *domain.User) error {
	user, err := s.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(id int, user *domain.User) (*domain.User, error) {
	user, err := s.userRepository.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUsers() ([]*domain.User, error) {
	users, err := s.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
