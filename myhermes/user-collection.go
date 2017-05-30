package main

import (
	"time"
)

type User struct {
	Id          int       `json:"id" hermes:"dbspace:users"`
	Birthdate   time.Time `json:"birthdate" hermes:"Type:time,editable,searchable,dbtype:date"`
	Firstname   string    `json:"firstname" hermes:'searchable,editable'`
	Lastname    string    `json:"lastname" hermes:"searchable,index,editable"`
	Email       string    `json:"email" validate"required,email" hermes:"searchable,index"`
	DisplayName string    `json:"displayname" hermes:"index,editable,searchable"`
	Gender      int       `json:"gender" hermes:"searchable,index,editable"`
	Is_Deleted  bool      `json:"is_deleted" hermes:"searchable"`
	Person      Person    `json:"person" db:"-" hermes:"one2one"`
	Person_Id   int       `json:"person_id"`
}
