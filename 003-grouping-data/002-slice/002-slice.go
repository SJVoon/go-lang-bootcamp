package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 8}

	//slicing
	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:3])

	//append
	fmt.Println(x)
	x = append(x, 77, 88, 99)
	fmt.Println(x)

	y := []int{123, 456, 789}
	x = append(x, y...)
	fmt.Println(x)

	//deleting
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
