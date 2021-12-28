package main

import (
	"fmt"
)

func main() {

	//switch i:=2+4;i{ //use of expressions
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5:
		fmt.Println("three four or five")
	default:
		fmt.Println("not one or two or four or five")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
	case i <= 20:
		fmt.Println("less that or equal to twenty")
	default:
		fmt.Println("greater thatn 20")

	}

	//type switch
	var z interface{} = 1.0
	switch z.(type) {
	case int:
		fmt.Println("z is an int")
	case float64:
		fmt.Println("z is float")
	}

}
