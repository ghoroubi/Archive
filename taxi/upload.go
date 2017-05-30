package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

type UploadRecord struct {
	Id  int
	Url string
}

func upload(c *gin.Context) {
	var path string
	uniquefid := uuid.NewV4().String()

	r := c.Request
	r.ParseMultipartForm(32 << 20)

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("error 31", err)
	}
	defer file.Close()

	url := "/home/Project/Go/src/taxi/" + path + uniquefid + ".jpg"
	out, err := os.Create(url)
	if err != nil {
		fmt.Println("error line 38", err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println("error line 43", err)
	}
	Imagemain, err := imaging.Open(url)
	if err != nil {
		fmt.Println("error in open image line 47", err)
	}
	w, h, err := getImage(url)
	if err != nil {
		fmt.Println("error in get w,h", err)
	}

	fmt.Println("Width", w)
	fmt.Println("Height", h)
	var cropCenter *image.NRGBA

	if w != h {
		if w > h {
			cropCenter = imaging.CropCenter(Imagemain, h, h)
			err = imaging.Save(cropCenter, url)
			if err != nil {
				fmt.Println("error in  crop image", err)
			}
		} else {
			cropCenter = imaging.CropCenter(Imagemain, w, w)
			err = imaging.Save(cropCenter, url)
			if err != nil {
				fmt.Println("error in crop image", err)
			}
		}
	}

	Imagemain, err = imaging.Open(url)
	if err != nil {
		fmt.Println("error in open image", err)
	}
	Image640 := imaging.Resize(Imagemain, 640, 640, imaging.Lanczos)
	err = imaging.Save(Image640, url)
	if err != nil {
		fmt.Println("error in save image", err)
	}

	f := &File{}
	f.File_Path = url
	f.Creation_Date = time.Now().Local()
	var id int
	err = db.QueryRow("insert into file(file_path,creation_date) values($1,$2) returning id", f.File_Path, f.Creation_Date).Scan(&id)
	fmt.Println("last inserted id", id)
	if err != nil {
		fmt.Println("error in last inserted id line 92", err)

	}
	uploadrecord := UploadRecord{}

	uploadrecord.Id = id
	uploadrecord.Url = url
	c.JSON(http.StatusOK, gin.H{
		"upload image": uploadrecord,
	})

}

func getImage(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("error in open line 108", err)
		return 0, 0, err
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("error in drcode line 114", err)
		return 0, 0, err

	}
	return image.Width, image.Height, nil
}
