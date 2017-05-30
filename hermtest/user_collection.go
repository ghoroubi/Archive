package main

import (
	"time"
)

type User struct {
	Id           int       `json:"id" hermes:"dbspace:users"`
	Birth_Date   time.Time `json:"birth_date" hermes:"type:time,editable,dbtype:date"`
	Display_Name string    `json:"display_name" hermes:"editable,searchable,index"`
	Email        string    `json:"email" validate:"required,email" hermes:"searchable,index"`
	Password     string    `json:"password,omitempty"`
	Is_Deleted   bool      `json:"is_deleted" hermes:"index,editable"`
	Gender       int       `json:"gender" hermes:"editable,index"`
	Person       Person    `json:"person" db:"-" hermes:"one2one"`
	Person_Id    int       `json:"person_id"`
}
