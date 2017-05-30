package main

type Person struct {
	Id        int     `json:"id" hermes:"dbspace:persons"`
	User_Name string  `json:"user_name" validate:"required"`
	Password  string  `json:"password,omitempty"`
	Skills    []Skill `json:"skills" db:"-" hermes:"one2many"`
	Needs     []Need  `json:"needs" db:"-" hermes:"one2many"`
}
