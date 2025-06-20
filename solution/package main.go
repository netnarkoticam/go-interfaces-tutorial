package main

import (
	"fmt"
	"math"
)

type Krug struct{
	Radius float64
}
type Kvadrat struct{
	Visota float64
	Shirirna float64
}
type Shape interface{
	Area() float64
}
func (k Krug) Area() float64{
	return math.Pi * k.Radius * k.Radius
}
func (k Kvadrat) Area() float64{
	return k.Visota * k.Shirirna
}
func main(){
krug:=Krug{5}
kvadrat:=Kvadrat{5, 6}
var shape1 Shape = krug
var shape2 Shape = kvadrat
fmt.Println("Площадь круга:", shape1.Area() )
fmt.Println("Площадь круга:", shape2.Area() )
}


