package services

import (
	"errors"
	"fmt"
	"notifications/models"
	"notifications/models/enums"
	"notifications/repositories"
	"time"
)

type UserService struct {
	userRepository         *repositories.UserRepository
	notificationRepository *repositories.NotificationRepository
}

func NewUserService(userRepository *repositories.UserRepository, notificationRepository *repositories.NotificationRepository) *UserService {
	return &UserService{
		userRepository:         userRepository,
		notificationRepository: notificationRepository,
	}
}

func (s *UserService) GetUsers() []models.User {
	if users := s.userRepository.GetUsers(); len(users) > 0 {
		return users
	}
	return nil
}

func (s *UserService) GetUserById(userId string) *models.User {
	user := s.userRepository.GetUserById(userId)
	if user != nil {
		return user
	}
	return nil
}

func (s *UserService) CreateUser(user models.User) {
	s.userRepository.CreateUser(user)
}

func (s *UserService) DeleteUserById(userId string) {
	s.userRepository.DeleteUserById(userId)
}

func (s *UserService) UserNotify(userId string, userNotification models.UserNotifyRequest) error {
	user := s.userRepository.GetUserById(userId)
	if user == nil {
		return errors.New(fmt.Sprintf("user_id: %s not found in db", userId))
	}

	notification := s.notificationRepository.GetNotificationById(userNotification.NotificationId)
	if notification == nil {
		return errors.New(fmt.Sprintf("notification_id: %s not found in db", userNotification.NotificationId))
	}

	if err := validateRestrictions(notification, user.Notifications); err != nil {
		return err
	}

	s.userRepository.AddUserNotification(userId, models.NewUserNotification(notification.Id, notification.Type))

	return nil
}

func validateRestrictions(notification *models.Notification, userNotifications []models.UserNotification) error {
	for _, restriction := range notification.Restrictions {
		from := time.Now().Add(getDuration(restriction.Timeframe.From, restriction.Timeframe.Scale))
		to := time.Now().Add(getDuration(restriction.Timeframe.To, restriction.Timeframe.Scale))
		notificationCount := 0
		for _, userNotification := range userNotifications {
			if (notification.Id == userNotification.Id || notification.Type == userNotification.Type) && (userNotification.DateCreated.After(from) && userNotification.DateCreated.Before(to)){
				notificationCount++
			}
		}
		if notificationCount >= restriction.Limit{
			return errors.New("restriction reached")
		}
	}
	return nil
}

func getDuration(quantity int, scale enums.TimeScale) time.Duration {
	duration := time.Hour
	switch scale {
	case enums.Hours:
		duration = time.Hour
	case enums.Days:
		duration = 24 * time.Hour
	case enums.Minutes:
		duration = time.Minute
	}
	return (time.Duration(int32(quantity)) * duration) * -1
}

