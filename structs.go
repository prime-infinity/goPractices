package main

import "fmt"

type customStruct struct {
	name  string
	items map[string]float64
	tizz  float64
}

//another struct
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func addDataToStruct(name string) customStruct {
	c := customStruct{
		name:  name,
		items: map[string]float64{},
		tizz:  0,
	}

	return c
}

/*
in the above function, we are returning a new customStruct type after initing it
*/

func main() {

	myStruct := addDataToStruct("osamede struct")
	fmt.Println(myStruct)
	fmt.Println(myStruct.name) //we can also get a particular item from it

	aDoctor := Doctor{
		number:    3,
		actorName: "JOhn pertwer",
		companions: []string{
			"Liz Shaw",
			"jo Grand",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[2])

	//anonymous structs
	aThis := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aThis)

	//structs are copied, and not referenced

}

/*
structs are used for custom data types, structs are great,
*/
