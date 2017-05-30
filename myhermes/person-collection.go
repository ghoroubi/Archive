package main

type Person struct {
	Id       int     `json:"id" hermes:"dbspace:persons"`
	UserName string  `json:"username" validate:"required" hermes:"index,searchable"`
	Password string  `json:"password,omitempty"`
	Skills   []Skill `json:"skills" db:"-" hermes:"one2many"`
	Needs    []Need  `json:"needs" db:"-" hermes:"one2many"`
}
