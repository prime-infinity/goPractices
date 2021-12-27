package main

import "fmt"

func main() {

	var n bool = true
	fmt.Printf("%v, %T", n, n)

	m := 1 == 1
	z := 1 == 2
	fmt.Printf("%v, %T", m, m)
	fmt.Printf("%v, %T", z, z)
	//in the above example, a boolean will be generated
	//from the result of the expression

	var t bool
	fmt.Printf("%v, %T", t, t)
	//in the above, the valueof t will be false cus in go, all values declared
	//until initialised, have a zero value, and for bool, zero value is false

	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))
	//in the above, go requires that you do a type conversion,

	/*
		also recall the bit operators
	*/
	/*fmt.Println(a & int(b))
	fmt.Println(a | int(b))
	fmt.Println(a ^ int(b))*/

	/*
		another intresting stuff is that go can use complex numbers as default first
		class citizens, observe
	*/

	var c complex128 = 1 + 2i
	c2 := 2 + 5.3i
	fmt.Println(c + c2)
	//we can also get the real and imaginary parts
	fmt.Printf("%v, %T", real(c), imag(c))
	//we can also convert normal numbers to complex numbers with real and img parts
	var c3 complex128 = complex(5, 12)
	fmt.Println(c3)

	s := "this is a string"
	l := []byte(s)
	fmt.Printf("%v, %T", l, l) //converting strings to byte
}
