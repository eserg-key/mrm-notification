package model

import "time"

type NotificationsInput struct {
	System        System        `json:"system"`
	Notifications Notifications `json:"notifications"`
}

func NewNotificationsInput(
	projectCode string,
	typeCode string,
	recipient []string,
	title string,
	message string,
	created time.Time,
	actions []NotificationsAction,
) *NotificationsInput {
	system := System{
		ProjectCode: projectCode,
		TypeCode:    typeCode,
		Recipient:   recipient,
	}

	notifications := Notifications{
		Title:   title,
		Message: message,
		Created: created,
		Actions: actions,
	}

	return &NotificationsInput{System: system, Notifications: notifications}
}

type System struct {
	ProjectCode string   `json:"project_code"`
	TypeCode    string   `json:"type_code"`
	Recipient   []string `json:"recipient"`
}

type Notifications struct {
	Title   string                `json:"title"`
	Message string                `json:"message"`
	Created time.Time             `json:"created"`
	Actions []NotificationsAction `json:"actions"`
}

type NotificationsAction struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Link  string `json:"link"`
}

func NewNotificationsAction(name string, color string, link string) *NotificationsAction {
	return &NotificationsAction{Name: name, Color: color, Link: link}
}
