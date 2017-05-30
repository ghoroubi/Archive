package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BarController struct {
}

var barcol BarCollection

func (col *BarController) GetVenuesWithLatitudeAndLongitude(c *gin.Context) {
	lat := c.Query("lat")
	lng := c.Query("lng")

	venues := barcol.GetVenuesWithLatitudeAndLongitude(lat, lng)
	// if err != nil {
	// 	HandleHttpError(c, err)
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"venues ": venues,
	})
}
