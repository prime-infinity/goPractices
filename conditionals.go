package main

import "fmt"

func main() {

	age := 45

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less that 40")
	} else { //catches all other conditions
		fmt.Println("age is not less that 45")
	}

	names := []string{"maria", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue //this is explained below
		}

		if index > 2 {
			fmt.Println("breaking at pos", index)
			break //explained below
		}

		fmt.Printf("the calue at pos %v is %v \n", index, value)
	}

}

/*
the continue keyword tells the program to break out of the current iteration of the loop
and begin a new iteration without getting past where the continue key was initially

the break keyword breaks out of the looop completely
*/
