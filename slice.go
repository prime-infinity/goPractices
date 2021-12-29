package main

import "fmt"

func main() {
	var scores = []int{232, 421, 44}
	//if the above was an array, it will
	//not change on line 12
	fmt.Println(scores)
	b := scores
	fmt.Println(b)
	b[1] = 5
	fmt.Println(scores)
	fmt.Println(b)

	//slices are reference tybes
	//arrays are value tybes
}
