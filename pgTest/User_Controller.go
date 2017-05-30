package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct{}

var usercol UserCollection

func (cont *UserController) CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(user)
	fmt.Printf("...", user)
	id, err := usercol.CreateUser(&user)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        "user created Succesfully",
		"last User ID:": id,
	})
}
func (cont *UserController) GetUsers(c *gin.Context) {

}
