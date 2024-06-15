package model

type Project struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Disabled bool   `json:"disabled"`
	Types    []Type `json:"types"`
}

func NewProject(name string, code string, disabled bool, types []Type) *Project {
	return &Project{Name: name, Code: code, Disabled: disabled, Types: types}
}

type Type struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Switchable bool   `json:"switchable"`
	Disabled   bool   `json:"disabled"`
}

func NewType(name string, code string, switchable bool, disabled bool) *Type {
	return &Type{Name: name, Code: code, Switchable: switchable, Disabled: disabled}
}
