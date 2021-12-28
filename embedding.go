package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   //embedd animal struct into bird
	SpeedKPH float32
	CanFly   bool
}

func main() {

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Austirlia"
	b.SpeedKPH = 43
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name) //we can also drill into other structs

	c := Bird{
		Animal:   Animal{Name: "Emuu", Origin: "Austirilia"},
		SpeedKPH: 34,
		CanFly:   false,
	}

	fmt.Println(c)

	/*
		here, we see that we can embedd on struct into another,
		giving the idea of inheritance.but not
	*/

}
