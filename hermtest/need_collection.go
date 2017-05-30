package main

type Need struct {
	Id        int    `json:"id" hermes:"dbspace:needs"`
	Title     string `json:"title" validate:"required"`
	Person_Id int    `json:"person_id" validate:"required"`
}
