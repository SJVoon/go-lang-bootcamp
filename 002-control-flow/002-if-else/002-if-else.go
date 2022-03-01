package main

import "fmt"

func main() {

	//normal if statement
	if true {
		fmt.Println("something")
	}

	//declaration in if statement
	if x := 42; x < 50 {
		fmt.Println("lesser")
	}

	y := 42
	if y == 40 {
		fmt.Println("40")
	} else if y == 41 {
		fmt.Println("41")
	} else {
		fmt.Println("42 more")
	}
}
