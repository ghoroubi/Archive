package main

type Skill struct {
	Id        int    `json:"id" hermes:"dbspace:skills"`
	Title     string `json:"title" validate:"required"`
	Person_Id int    `json:"person_id" validate:"required"`
}
