package main

import (
	"fmt"
)

func main() {
	x := 42
	y := 43
	z := " Hello"

	//U Hello
	fmt.Println(string(x+y) + z)
	//85 Hello
	fmt.Println(fmt.Sprint(x+y) + z)

	foo()
	bar()
}

/*
	scope of variable understanding
*/

//declare variable with var
var a = 12

//ASSIGNS the ZERO VALUE of TYPE int to "b"
//false for booleans, 0 for integers, 0.0 for floats, "" for strings,
//and nil for pointers, functions, interfaces, slices, channels, and maps.
var b int

////error, syntax error: non-declaration statement outside function body
// short declaration can only at package level
// c := 5

////error, syntax error: unexpected newline, expecting type
// var d

func foo() {
	//12
	fmt.Println(a)
	//0
	fmt.Println(b)
	b = 13
	//13
	fmt.Println(b)
	c := 14
	//14
	fmt.Println(c)
}

func bar() {
	//12
	fmt.Println(a)
	//13
	fmt.Println(b)

	////error, undefined: c
	// fmt.Println(c)
}
