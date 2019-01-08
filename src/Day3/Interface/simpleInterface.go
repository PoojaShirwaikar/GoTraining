package main

import "fmt"

//define interafce

type Shape interface {
	Area()
	Perimeter()
}

type Triangle struct {
	side int
}

type Rectangle struct {
	side int
}

func (tra Triangle) Area() {
	fmt.Println("area of triangle")
}
func (tra Triangle) Perimeter() {
	fmt.Println("Perimeter")

}

func Calculate(shape Shape) {
	fmt.Println("calculate is called")
	shape.Area()
	shape.Perimeter()
}

func main() {
	trac1 := Triangle{20}
	//trac1.Area()

	Calculate(trac1)

}
