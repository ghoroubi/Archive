package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)
var dbObject *sqlx.DB
var uController UserController
var uCollection UserCollection
var schema = `CREATE IF NOT EXIST users(
ID SERIAL PRIMARY KEY NOT NULL,
DisplayName text,
Email text UNIQUE NOT NULL,
Password text NOT NULL
);`
///////////////////////////////////
type UserCollection struct{}
type UserController struct{}
type User struct{
	ID int `jason:"id"`
	DisplayName string `json:"display_name"`
	Email string `json:"email"`
	Password string	 `json:"password"`
	}
//////////////////////////////////
func (usController *UserController)CreateUser( c *gin.Context) error{
var user User
	c.BindJSON(user)
	err:=uCollection.CreateUser(&user)
	if err!=nil {
		log.Panic(err)
		return err
	}
	return nil
}
//////////////////////////////////
func (usCollection *UserCollection)CreateUser(user *User) error {
var err error
	result:=fmt.Sprintf("insert into users(DisplayName,Email,Password) values(%s,%s,%s)",
		user.DisplayName,user.Email,user.Password)
	dbObject.QueryRow(result)
	if err!=nil {
		log.Panic(err)
		return  err
	}

	return nil
}
///////////////////////////////////
//noinspection GoTypesCompatibility
func main(){
	var uc UserController
	var err error
	dbObject , _ =sqlx.Connect("postgres","user=postgres password=123456 dbname=dbTest sslmode=disable")
	if err!=nil {
		log.Panic(err)
	}
	router:=gin.Default()

	dbObject.MustExec(schema)
	router.POST("/CreateUser",uc.CreateUser)
	router.Run(":8585")
}