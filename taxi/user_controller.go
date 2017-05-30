package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var usercol UserCollection

func (col *UserController) Create(c *gin.Context) {

	phone_number := c.Param("phone")
	err, result := usercol.Create(phone_number)
	if err != nil {
		HandleHttpError(c, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"create user with the phone number is": phone_number,
		" and id is ":                          result,
	})

}

func (col *UserController) Login(c *gin.Context) {
	user := Users{}
	user.AgentToken.Device.Ip = c.ClientIP()
	err := c.BindJSON(&user)
	fmt.Println("user in ", user)
	if err != nil {
		fmt.Println("error in bindjson", err)

	}

	result, err := usercol.Login(&user)
	if err != nil {
		fmt.Println("error in return value in login", err)

	} else {

		c.JSON(http.StatusOK, gin.H{
			"login user": result,
		})
	}
}
func (col *UserController) Logout(c *gin.Context) {

	user := Users{}

	user.AgentToken.Creation_Date = time.Now()

	err := c.BindJSON(&user)
	if err != nil {

		HandleHttpError(c, err)
	}
	result, err := usercol.Logout(&user)

	if err != nil {
		HandleHttpError(c, err)
	} else {

		c.JSON(http.StatusOK, result)
	}
}

func (col *UserController) EditProfile(c *gin.Context) {

	user := Users{}
	err := c.BindJSON(&user)
	if err != nil {

		fmt.Println("error in edit profile is :", err)
	}

	result, err := usercol.EditProfile(&user)

	if err != nil {

		fmt.Println("error in return value in editprofile is", err)
	}
	c.JSON(http.StatusOK, result)

}

func (col *UserController) List(c *gin.Context) {
	pg, err := ReadPaging(c)
	fmt.Println("paging", pg)
	page := pg.Num
	fmt.Println("page", page)
	pageSize := pg.Size
	fmt.Println("pagesize", pageSize)
	token := c.Query("token")
	fmt.Println("token", token)
	result, err := usercol.List(page, pageSize)
	if err != nil {
		HandleHttpError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
