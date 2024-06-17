package model

type Project struct {
	Name     string `json:"name"`     // Название проекта
	Code     string `json:"code"`     // Код проекта
	Disabled bool   `json:"disabled"` // Статус работы проекта true|false
	Types    []Type `json:"types"`    // Типы уведомлений
}

func NewProject(name string, code string, disabled bool, types []Type) Project {
	return Project{Name: name, Code: code, Disabled: disabled, Types: types}
}

type Type struct {
	Name       string `json:"name"`       // Название уведомления
	Code       string `json:"code"`       // Код уведомлений
	Switchable bool   `json:"switchable"` // Переключатель в профиле true - отображается в настройках
	Disabled   bool   `json:"disabled"`   //  Статус работы уведомления true|false
}

func NewType(name string, code string, switchable bool, disabled bool) Type {
	return Type{Name: name, Code: code, Switchable: switchable, Disabled: disabled}
}
