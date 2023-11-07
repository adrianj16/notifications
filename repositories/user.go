package repositories

import (
	"notifications/models"
)

type UserRepository struct {
	Users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Users: []models.User{},
	}
}

func (r *UserRepository) GetUserByUserId(userId string) *models.User {
	for _, user := range r.Users {
		if user.UserId == userId{
			return &user
		}
	}
	return nil
}

func (r *UserRepository) AddUser(user models.User) error {
	r.Users = append(r.Users, user)
	return nil
}
