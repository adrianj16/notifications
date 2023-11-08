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

func (r *UserRepository) GetUsers() []models.User {
	return r.Users
}

func (r *UserRepository) GetUserById(userId string) *models.User {
	for _, user := range r.Users {
		if user.UserId == userId{
			return &user
		}
	}
	return nil
}

func (r *UserRepository) CreateUser(user models.User) {
	r.Users = append(r.Users, user)
}

func (r *UserRepository) DeleteUserById(userId string) {
	foundIndex := -1
	for i, user := range r.Users {
		if user.UserId == userId{
			foundIndex = i
		}
	}
	if foundIndex != -1 {
		r.Users = append(r.Users[:foundIndex], r.Users[foundIndex+1:]...)
	}
}

func (r *UserRepository) AddUserNotification(userId string, notification models.UserNotification) {
	for i, user := range r.Users {
		if user.UserId == userId{
			r.Users[i].Notifications = append(r.Users[i].Notifications, notification)
		}
	}
}
