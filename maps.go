package main

import "fmt"

func main() {

	statePopulations := map[string]int{ //we are mapping one datatype to another
		"califonia": 232434352,
		"Texas":     23232323243,
		"Florida":   223423343,
		"New york":  23423234,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Texas"]) //pull out a particular value using its key

	statePopulations["Gorgia"] = 23434343 //we can add items like this
	fmt.Println(statePopulations)

	delete(statePopulations, "Texas") //we can delete items like this

	pop, ok := statePopulations["Texas"] //this will properly interogate the map, we can also use the _
	fmt.Println(pop, ok)

	fmt.Println(len(statePopulations)) //length of map

	//declare a map without initing it
	//statePopulations := make(map[string]int)
	//statePopulations = map[string]int{}

	sp := statePopulations
	delete(sp, "New york")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	/*
	 note, in the above, maps are passed by reference, contuary to arrays and slices,
	 this is like shallow copying in javascript, means that they both point to same data
	 changing one will mutate other
	*/

}

/*
a map is like an object in javascript, but a little bit different
it is like a very strict object. it has key/value pairs and they
must all be of the same type

the key types in maps must be able to be tested for equality, this includes
all data types except slices,other maps, and functions. note, an array can however be used

note:the return order of a map is not guranteed
*/
