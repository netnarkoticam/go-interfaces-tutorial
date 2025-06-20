package main

import "fmt"

type Vehicle interface {
	Start()
	Stop()
}

type Car struct{
	Name string
}

type Moto struct{
	Name string
}

func (c Car) Start(){
	fmt.Println("Машина модели", c.Name, "начала движение" )
}

func (c Car) Stop(){
	fmt.Println("Машина модели", c.Name, "закончила движение" )
}
func (m Moto) Start(){
	fmt.Println("Мотоцикл модели", m.Name, "начал движение" )
}
func (m Moto) Stop(){
	fmt.Println("Мотоцикл модели", m.Name, "закончил движение" )
}

func main(){
	Mashina := Car{Name: "Камрюха"}
	
	Motik := Moto{Name: "Харлей"}

	Mashina.Start()
	Mashina.Stop()
	Motik.Start()
	Motik.Stop()
	
}