package models

type User struct {
	UserId        string             `json:"user_id"`
	UserName      string             `json:"user_name"`
	Type          string             `json:"type"`
	Email         string             `json:"email"`
	Phone         string             `json:"phone"`
	Notifications []UserNotification `json:"notifications"`
}

type UserNotification struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	DateCreated string `json:"date_created"`
}
