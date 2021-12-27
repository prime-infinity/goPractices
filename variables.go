package main

import "fmt"

//we can also do the below to group variables
var (
	actorName    string = "Elisabeth Sladed"
	companion    string = "Sarah jane smith"
	doctorNumber int    = 3
	season       int    = 11
)

var I int = 44 //this will be visible to outside pageages cus of the uppercase

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

	fmt.Println(actorName)

	//we can also do type conversion from one type to another,
	//observe
	var z int = 34
	var j float32
	j = float32(z) //z has been converted to float32 and assigned to j
	fmt.Println(j)

}

/*
variables cannot be declared twice in the same scope,
the vraiable with the innermost scope will take presedence


also remember that using uppercase var names at package level will make
the variables visible to outside packages
*/
