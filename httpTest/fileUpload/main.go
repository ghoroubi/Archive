package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"io"
	//"log"
	//"os"
)

// func main() {
// 	r := gin.Default()
// 	filepath := "/home/upload/mom.jpg"
// 	r.POST("/upload", func(c *gin.Context) {
// 		file, header, err := c.Request.FormFile(filepath)
// 		filename := header.Filename
// 		fmt.Println(filename)
// 		out, err := os.Create("/home/" + filename + ".png")
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		defer out.Close()
// 		_, err = io.Copy(out, file)
// 		if err != nil {
// 			panic(err)
// 		}
// 		r.Run(":8585")
// 	})

// }
func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {

		_, header, err := c.Request.FormFile("upload")
		if err != nil {
			//panic(err)
		}
		filename := header.Filename
		fmt.Println(filename)
		// out, err := os.Create("/home/ghoroubi/" + filename + ".png")
		// if err != nil {
		// 	//	log.Fatal(err)
		// }
		// defer out.Close()
		// _, err = io.Copy(out, file)
		// if err != nil {
		// 	//log.Fatal(err)
		// }
	})
	router.Run(":8080")
}
