package services

import (
	"errors"
	"notifications/models"
	"notifications/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) GetUserById(userId string) *models.User {
	user := s.userRepository.GetUserByUserId(userId)
	if user != nil {
		return user
	}
	return nil
}

func (s *UserService) Register(userId string) error {
	user := s.userRepository.GetUserByUserId(userId)
	if user != nil {
		return nil
	}
	return errors.New("user not exist")
}