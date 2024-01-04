package services

import (
	dtos "notification-system/internal/dtos/user"
	"notification-system/internal/entitys"
	"notification-system/internal/repositories"
)

type IUserService interface {
	Create(input dtos.CreateUserDTO) (*dtos.CreateMessage, error)
}

type UserService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user dtos.CreateUserDTO) (*dtos.CreateMessage, error) {
	users := &entitys.Users{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CompanyID: user.CompanyID,
		GroupID:   user.GroupID,
	}

	_, err := s.repo.Create(users)

	if err != nil {
		return nil, err
	}

	return &dtos.CreateMessage{Message: "User created successfully"}, nil
}
