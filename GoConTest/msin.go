package main

import (
	"fmt"
	"math"
	//"github.com/satori/go.uuid"
	//"strings"
	//"net/http"
)

type gPoint struct {
	x, y, z float64
}

type iPoint interface {
	distance() float64
	getPow(int) int
}

func (gp gPoint) getPow(x int) int {
	return x * int(gp.x)
}

func (gp1 gPoint) distance() float64 {
	s := math.Sqrt(math.Pow((gp1.x-2.0), 2) + math.Pow((gp1.y-2.0), 2) + math.Pow((gp1.z-2.0), 2))
	return s
}

func main() {
	var p iPoint
	p = gPoint{5, 3, 4}
	fmt.Println(p.getPow(5))
	fmt.Printf("%v  %t", p, p)
}
