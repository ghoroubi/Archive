package main1

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main1() {
	code := http.StatusOK
	r := gin.Default()
	// r.GET("/user/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.String(code, "your ID is:%s", id)
	// })

	r.GET("/", func(c *gin.Context) {
		//firstname := c.DefaultQuery("firstname", "Dear Guest")
		//lastname := c.Query("lastname")
		//c.String(code, "Welcome , %s  %s", firstname, lastname)
		c.String(code, "%s", "<html> <head> <title> By Nima Ghoroubi</title> <body> <H1>Test </body> </html>")
	})
	r.Run(":80")
	//fmt.Println("Test")
}
