package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type FileController struct {
// }

//var filcol FileCollection

// func (col *FileController) UploadImage(c *gin.Context) {

// 	contentType := c.PostForm("typecontent")

// 	obj, err := filcol.UploadImage(url, contentType) // obj.url ,obj.id
// 	if err != nil {
// 		fmt.Println("error in file", err)
// 	}
// 	c.JSON(http.StatusOK, obj)
// }

// 	pg, err := ReadPaging(c)
// 	fmt.Println("paging", pg)
// 	page := pg.Num
// 	fmt.Println("page", page)
// 	page_Size := pg.Size
// 	fmt.Println("pagesize", page_Size)
// 	token := c.Query("token")
// 	fmt.Println("token", token)
// 	result, err := filcol.List(page, page_Size, token)
// 	if err != nil {
// 		HandleHttpError(c, err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{"result": result})
