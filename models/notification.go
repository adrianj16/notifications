package models

import "notifications/models/enums"

type Notification struct {
	Id           string                    `json:"id"`
	Type         string                    `json:"type"`
	Message      string                    `json:"message"`
	Method       string                    `json:"method"`
	Data         []string                  `json:"data"`
	Restrictions []NotificationRestriction `json:"restrictions"`
}

type NotificationRestriction struct {
	Limit     int       `json:"quantity"`
	Timeframe Timeframe `json:"timeframe"`
}

type Timeframe struct {
	From  int       `json:"from"`
	To    int             `json:"to"`
	Scale enums.TimeScale `json:"scale"`
}

type UserNotifyRequest struct {
	NotificationId string                 `json:"notification_id"`
	Data           map[string]interface{} `json:"data"`
}