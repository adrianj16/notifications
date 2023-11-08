package repositories

import (
	"notifications/models"
)

type NotificationRepository struct {
	Notifications []models.Notification
}

func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{
		Notifications: []models.Notification{},
	}
}

func (r *NotificationRepository) GetNotifications() []models.Notification {
	return r.Notifications
}

func (r *NotificationRepository) GetNotificationById(notificationId string) *models.Notification {
	for _, notification := range r.Notifications {
		if notification.Id == notificationId {
			return &notification
		}
	}
	return nil
}

func (r *NotificationRepository) CreateNotification(notification models.Notification) {
	r.Notifications = append(r.Notifications, notification)
}

func (r *NotificationRepository) DeleteNotificationById(notificationId string) {
	foundIndex := -1
	for i, notification := range r.Notifications {
		if notification.Id == notificationId {
			foundIndex = i
		}
	}
	if foundIndex != -1 {
		r.Notifications = append(r.Notifications[:foundIndex], r.Notifications[foundIndex+1:]...)
	}
}

func (r *NotificationRepository) AddNotificationRestriction(notificationId string, restriction models.NotificationRestriction) {
	for i, notification := range r.Notifications {
		if notification.Id == notificationId {
			r.Notifications[i].Restrictions = append(r.Notifications[i].Restrictions, restriction)
		}
	}
}