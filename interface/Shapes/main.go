package main

import (
	"fmt"
)

type Shapes interface {
	Area() int
	Perimeter() int
}
type Square struct {
	length    int
	area      int
	perimeter int
}
type Rectangle struct {
	length    int
	width     int
	area      int
	perimeter int
}

func (s Square) Area() int {
	s.area = s.length * s.length
	return s.area
}
func (r Rectangle) Area() int {
	r.area = r.length * r.width
	return r.area
}
func (s Square) Perimeter() int {
	s.perimeter = (s.length) * 4
	return s.perimeter
}
func (r Rectangle) Perimeter() int {
	r.perimeter = (r.length + r.width) * 2
	return r.perimeter
}

/////////////////////////////////////
func main() {
	var iSquare, iRectangle Shapes
	var square Square = Square{5, 0, 0}
	var rectangle Rectangle = Rectangle{6, 8, 0, 0}
	iSquare = square
	iRectangle = rectangle

	fmt.Printf("Area of Square=%d \n", iSquare.Area())
	fmt.Printf("Perimeter of Square=%d \n", iSquare.Perimeter())
	fmt.Println("***********************************")
	fmt.Printf("Area of Rectangle=%d \n", iRectangle.Area())
	fmt.Printf("Perimeter of Rectangle=%d \n", iRectangle.Perimeter())
}
