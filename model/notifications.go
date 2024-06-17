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
) NotificationsInput {
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

	return NotificationsInput{System: system, Notifications: notifications}
}

type System struct {
	ProjectCode string   `json:"project_code"` // Код зарегистрированного проекта
	TypeCode    string   `json:"type_code"`    //  Код типа уведомления
	Recipient   []string `json:"recipient"`    // Получатели
}

type Notifications struct {
	Title   string                `json:"title"`   // Заголовок сообщения
	Message string                `json:"message"` // Текст сообщения
	Created time.Time             `json:"created"` // Дата
	Actions []NotificationsAction `json:"actions"` // Кнопки
}

type NotificationsAction struct {
	Name  string `json:"name"`  // Текст кнопки
	Color string `json:"color"` // Цвет кнопки primary|secondary
	Link  string `json:"link"`  // Ссылка кнопки
}

func NewNotificationsAction(name string, color string, link string) NotificationsAction {
	return NotificationsAction{Name: name, Color: color, Link: link}
}
