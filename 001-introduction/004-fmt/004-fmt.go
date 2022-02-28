package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	// int *int
	fmt.Printf("%T %T\n", y, &y)
	// 42 42 2a 52 101010
	fmt.Printf("%v %d %x %o %b\n", y, y, y, y, y)

	z := fmt.Sprintf("%b\n", y)
	fmt.Print(z)
}
