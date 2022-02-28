package main

import "fmt"

// declare my own type name mytype with underlying type int
type mytype int

// assign mytype to a new var x
var x mytype

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	//error, cannot use x (type mytype) as type int in assignment
	// y = x
	y = int(x)
	fmt.Println(y)
}
