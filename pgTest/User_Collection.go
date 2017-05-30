package main

import (
	//"github.com/gin-gonic/gin"
	// "time"
	"crypto/rand"
	"fmt"
	"io"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID          int `json:"id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}
type UserCollection struct{}


func (col *UserCollection) CreateUser(user *User) (int, error) {
	var id int

	fmt.Println()
	res := fmt.Sprintf("insert into users() values(%s,%s,%s,%s)", user.UserName, user.Password, user.Email, user.DisplayName)
	err := DB.QueryRow(res).Scan(&id)

	fmt.Printf("Last Inserted Id : %d", id)
	if err != nil {
		return 0, nil
	}
	return id, nil
}
func (col *UserCollection) GetUsers(c *gin.Context)(error){
var users []User
	query:="SELECT * FROM users order by ID"
	err:=DB.Select(&users,query)
	if err==nil {
	c.JSON(200,&users)
	}

return nil

}
