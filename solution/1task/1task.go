
package main

import (
	"fmt"
	"math"
)


type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64 
}

type Rectangle struct {
	Width  float64 
	Height float64 
}


func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}


func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 5, Height: 5}

	
	var shape1 Shape = circle
	var shape2 Shape = rectangle

	fmt.Println("Площадь круга:", shape1.Area())
	fmt.Println("Площадь прямоугольника:", shape2.Area()) 
}


