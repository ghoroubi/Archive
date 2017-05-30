package main

type Skill struct {
	Id        int    `json:"id" hermes:"dbspace:skills"`
	Title     string `json:"title" hermes:"editable,searchable,index"`
	Person_Id int    `json:"person_id" validate:"required"`
}
