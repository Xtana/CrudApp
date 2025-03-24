package service

import (
	"crudapp/internal/domain"
	"crudapp/internal/repository"
)

type UserService interface {
	CreateUser(user *domain.User) 	error
	GetUserById(id int64) 			(*domain.User, error)
	GetAllUsers() 					([]*domain.User, error)
	UpdateUser(user *domain.User) 	error
	DeleteUser(id int64) 			error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUserById(id int64) (*domain.User, error) {
	return s.repo.GetById(id)
}

func (s *userService) GetAllUsers() ([]*domain.User, error) {
	return s.repo.GetAll()
}

func (s *userService) UpdateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *userService) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}