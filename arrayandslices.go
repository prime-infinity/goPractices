package main

import "fmt"

func main() {

	//arrays
	var ages [3]int = [3]int{20, 31, 43}
	fmt.Println(ages)

	grades := [3]int{2, 32, 44}
	//graded := [...]int{3,4,2,43} //this will also work, no need to declare size of array
	grades[2] = 33
	fmt.Println(grades)

	//len(grades) //to get length of array

	/*
		a very intrestin tin about arrays in go is that they are values themselves
		so if we copy an array, we are not copying the reference to the underlying values,
		we are instead copying the array itself to another value, this is similar to deepcopying in
		javascript. observe
	*/

	a := [...]int{1, 2, 3}
	b := a //here, a is copied into a
	b[1] = 4
	fmt.Println(a)
	fmt.Println(b)
	//inthe above, it will be observed that b is its own arry, so it does not change a
	//this means that go is copying the full array, this might be cause for concern when passing large arrays
	//into functions

	//slices
	var scores = []int{232, 421, 44} //this creates a slice
	scores[2] = 32
	scores = append(scores, 53) //add element. actually returns new slice

	//slice ranges
	rangeOne := scores[1:3]  //go from position 1 to 3, not including 3
	rangeTwo := scores[2:]   //gfrom 3 to end, including end
	rangeThree := scores[:3] //gfrm start to 3, not inculuding 3

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	removedMiddle := append(scores[:2], scores[3:]...) //remove element from middle of array
	fmt.Println(removedMiddle)

}

/*
arrays in go are kinda too strict, they have to always be given
a fixed length at compile time

*/
