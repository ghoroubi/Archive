package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
  CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    gender integer NOT NULL DEFAULT 0,
    birth_date timestamp with time zone NOT NULL DEFAULT '0001-01-01 03:25:44+03:25:44',
    lng_init  real NOT NULL DEFAULT 0,
    lat_init real NOT NULL DEFAULT 0,
    lat_current real NOT NULL DEFAULT 0,
    lng_current real NOT NULL DEFAULT 0,
    registration_date  timestamp with time zone NOT NULL DEFAULT '0001-01-01 03:25:44+03:25:44',
    image_url text NOT NULL DEFAULT '' 
    
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
)`
var db *sqlx.DB

func main() {
	InitMessages()
	var usercon UserController
	var barcol BarController
	var err error
	router := gin.Default()
	db, err = sqlx.Connect("postgres", "user=postgres password=123456 dbname=barsee sslmode=disable")

	if err != nil {

		log.Fatalln(err)
	}
	db.MustExec(schema)

	router.POST("/Create_User", usercon.Create_User)
	router.GET("/user", usercon.Get_User)
	router.POST("/UpdateLocation", usercon.UpdateLocation)
	router.POST("/Logout", usercon.Logout)
	router.GET("/GetVenue", barcol.GetVenuesWithLatitudeAndLongitude)

	router.Run(":8080")
}
