package models

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
	From  string `json:"from"`
	To    string `json:"to"`
	Scale string `json:"scale"`
}
