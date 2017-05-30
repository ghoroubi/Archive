package main

type Favorite struct {
	Id        int     `json:"id" hermes:"dbspace:favorite"`
	User_Id   int     `json:"user_id" validate:"required"`
	Longitude float32 `json:"longitude" validate:"required"`
	Latitude  float32 `json:"latitude" validate:"required"`
	Title     string  `json:"title" validate:"required"`
}
