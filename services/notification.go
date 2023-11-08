package services

import (
	"notifications/models"
	"notifications/repositories"
)

type NotificationService struct {
	notificationRepository *repositories.NotificationRepository
}

func NewNotificationService(notificationRepository *repositories.NotificationRepository) *NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepository,
	}
}

func (s *NotificationService) GetNotifications() []models.Notification {
	if notifications := s.notificationRepository.GetNotifications(); len(notifications) > 0 {
		return notifications
	}
	return nil
}

func (s *NotificationService) GetNotificationById(notificationId string) *models.Notification {
	notification := s.notificationRepository.GetNotificationById(notificationId)
	if notification != nil {
		return notification
	}
	return nil
}

func (s *NotificationService) CreateNotification(notification models.Notification) {
	s.notificationRepository.CreateNotification(notification)
}

func (s *NotificationService) DeleteNotificationById(notificationId string) {
	s.notificationRepository.DeleteNotificationById(notificationId)
}

func (s *NotificationService) AddNotificationRestriction(notificationId string, restriction models.NotificationRestriction) {
	s.notificationRepository.AddNotificationRestriction(notificationId, restriction)
}
