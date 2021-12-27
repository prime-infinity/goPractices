package main

import "fmt"

func main() {

	//string in go are double quoted
	var nameOne string = "name One"
	var nameTwo = "name Two"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Surname"
	nameThree = "Isset"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "name four" //this method cannot be used outside of a function
	fmt.Println(nameFour)

	//ints
	var ageOne int = 23
	var ageTwo = 40
	ageThree := 223
	fmt.Println(ageOne, ageTwo, ageThree)

	//an unsigned int, uint, cannnot have a negative number
	var numberr uint = 34
	fmt.Println(numberr)

	//floats are numbers with decimal points
	var scoreOne float32 = 23.53 //the bit size must always be infered
	var scoreTwo float64 = 4334.23232
	scoreThree := 434.12
	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
