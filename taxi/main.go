package main

import (
	//"database/sql"
	//"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
  CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    phone_number text NOT NULL ,
    firstname text NOT NULL DEFAULT '' ,
    lastname text NOT NULL DEFAULT '',
    email text NOT NULL DEFAULT '',
    lng  real  NOT NULL DEFAULT 0 ,
    lat  real   NOT NULL DEFAULT 0 ,
    registration_date  timestamp with time zone NOT NULL DEFAULT '0001-01-01 03:25:44+03:25:44',
    image_url text NOT NULL DEFAULT '' ,
    is_spammed bool NOT NULL DEFAULT FALSE,
    is_active bool NOT NULL DEFAULT FALSE

);
CREATE TABLE IF NOT EXISTS agent_tokens (
    id SERIAL PRIMARY KEY NOT NULL,
	user_id integer REFERENCES users(id),
	token text NOT NULL DEFAULT '',
	creation_date timestamp with time zone  NOT NULL DEFAULT '0001-01-01 03:25:44+03:25:44',
    is_expire bool NOT NULL DEFAULT FALSE,
    type text NOT NULL DEFAULT ''
);
CREATE TABLE IF NOT EXISTS device (
    id SERIAL PRIMARY KEY NOT NULL,
	usertokens_id integer REFERENCES agent_tokens(id),
	model text NOT NULL DEFAULT '',
    platform text NOT NULL DEFAULT '',
    uuid text NOT NULL DEFAULT '',
    version text NOT NULL DEFAULT '',
    ip text NOT NULL DEFAULT '',
    cm_id text NOT NULL DEFAULT ''
);
CREATE TABLE IF NOT EXISTS passenger (
    id SERIAL PRIMARY KEY NOT NULL,
	user_id integer REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS drivers (
    id SERIAL PRIMARY KEY NOT NULL,
	user_id integer REFERENCES users(id),
	car_model integer NOT NULL DEFAULT 0 ,
	car_license_plate text NOT NULL DEFAULT '',
	car_year integer NOT NULL DEFAULT 0,
	car_color integer NOT NULL DEFAULT 0
);
CREATE TABLE IF NOT EXISTS trips (
    id SERIAL PRIMARY KEY NOT NULL,
	passenger_id integer REFERENCES passenger(id),
	driver_id integer REFERENCES drivers(id),
	origin_longitude  real  NOT NULL DEFAULT 0 ,
    origin_latitude  real   NOT NULL DEFAULT 0 ,
    source_longitude real   NOT NULL DEFAULT 0,
    source_latitude real   NOT NULL DEFAULT 0,
    status integer NOT NULL DEFAULT 0 
);
CREATE TABLE IF NOT EXISTS offer (
    id SERIAL PRIMARY KEY NOT NULL,
	trip_id integer REFERENCES trips(id),
	driver_id integer REFERENCES drivers(id),
	status integer NOT NULL DEFAULT 0 
);
CREATE TABLE IF NOT EXISTS car_model (
    id SERIAL PRIMARY KEY NOT NULL,
	title text NOT NULL DEFAULT ''
);
CREATE TABLE IF NOT EXISTS car_color (
    id SERIAL PRIMARY KEY NOT NULL,
	title text NOT NULL DEFAULT ''
);
CREATE TABLE IF NOT EXISTS favorite (
    id SERIAL PRIMARY KEY NOT NULL,
	user_id integer REFERENCES users(id),
	longitude  real  NOT NULL DEFAULT 0 ,
	latitude real  NOT NULL DEFAULT 0 ,
	title text NOT NULL DEFAULT ''
);
CREATE TABLE IF NOT EXISTS agency(
    id SERIAL PRIMARY KEY NOT NULL,
    driver_id integer REFERENCES drivers(id)
);
CREATE TABLE IF NOT EXISTS file(
    id SERIAL PRIMARY KEY NOT NULL,
    file_path text NOT NULL DEFAULT '',
    creation_date timestamp with time zone NOT NULL DEFAULT '0001-01-01 03:25:44+03:25:44',
    is_deleted bool NOT NULL DEFAULT FALSE
)`

var db *sqlx.DB

func main() {
	InitMessages()
	var usercon UserController
	//var filcon FileController
	router := gin.Default()
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres password=123456 dbname=yabi sslmode=disable")

	if err != nil {

		log.Fatalln(err)
	}
	db.MustExec(schema)

	router.POST("/Create/:phone", usercon.Create)
	router.POST("/Login", usercon.Login)
	router.POST("/Logout", usercon.Logout)
	router.POST("/List", usercon.List)
	//router.POST("/Url", filcon.Geturl)
	router.POST("/Upload", upload)
	router.POST("/Upload_file", Upload)
	router.POST("/EditProfile", usercon.EditProfile)
	router.Run(":8080")
}
