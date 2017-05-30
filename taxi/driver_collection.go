package main

//todo agency should be required
type Driver struct {
	Id                int    `json:"id"`
	User_Id           int    `json:"user_id"`
	Agency_Id         int    `json:"agency_id"`
	Car_Model         int    `json:"car_model" validate:"required" `
	Car_License_Plate string `json:"car_license_plate" validate:"required" `
	Car_Year          int    `json:"car_year" validate:"required" `
	Car_Color         int    `json:"car_color" validate:"required" `
}
