package main

import "fmt"

func main() {

	x := 0
	for x < 5 { //while x is < than 5
		//fmt.Println("the value of x is :", x)
		x++
	}

	for i := 0; i < 5; i++ { //the tradtional for loop
		//fmt.Println("the value of x is :", i)
	}

	names := []string{"Akonofua", "osamede", "omokhomion", "prime"}
	for index, value := range names { //looping over a collection. this can also be done with atrad for loop
		fmt.Println("the index is", index, "the value is", value)

		//value = "something" //this will not work
	}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//we can also loop more than one stuff,
	//the below code did not work, you must go and search it out yourself
	/*for l, j := 0, 0; l < 5; j = l + 1, j + 1 {

		fmt.Println(l)

	}*/

}

//in go, the for keyword at the time of this, the only kind of loop available
