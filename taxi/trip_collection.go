package main

//"database/sql"

type Trip struct {
	Id               int     `json:"id"`
	Origin_Longitude float32 `json:"origin_longitude" validate:"required"`
	Origin_Latitude  float32 `json:"origin_latitude" validate:"required"`
	Source_Longitude float32 `json:"source_longitude" validate:"required"`
	Source_Latitude  float32 `json:"source_latitude" validate:"required"`
	Status           int     `json:"status"`
	Driver_Id        int     `json:"driver_id"`
	Passenger_Id     int     `json:"passenger_id"`
}
type TripCollection struct {
}

// func (col *TripCollection) Create(token string) ([]Trip, int) {

// 	var trip []Trip

// }
