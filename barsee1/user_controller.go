package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var usercol UserCollection

func (col *UserController) Create_User(c *gin.Context) {
	var user User
	user.Device.Ip = c.ClientIP()
	c.BindJSON(&user)

	fmt.Println("...", user)
	token, id, err := usercol.Create_User(&user)
	if err != nil {
		HandleHttpError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token user is ": token,
		"id user is":     id,
	})
}
func (col *UserController) Logout(c *gin.Context) {

	user := User{}
	err := c.BindJSON(&user)
	if err != nil {

		HandleHttpError(c, err)
	}
	text, err := usercol.Logout(&user)

	if err != nil {
		HandleHttpError(c, err)
	} else {

		c.JSON(http.StatusOK, text)
	}
}
func (col *UserController) UpdateLocation(c *gin.Context) {
	var err error

	var text string
	user := User{}
	err = c.BindJSON(&user)
	if err != nil {

		HandleHttpError(c, err)
	}

	text, err = usercol.UpdateLocation(&user)

	if err != nil {
		HandleHttpError(c, err)
	} else {

		c.JSON(http.StatusOK, text)
	}
}
func (col *UserController) Get_User(c *gin.Context) {
	var err error
	id := c.Query("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("err", err)
	}
	user := User{}
	user, err = usercol.Get_User(i)
	fmt.Println("error ", err)
	if err != nil {
		HandleHttpError(c, err)
	} else {

		c.JSON(http.StatusOK, user)
	}
}
