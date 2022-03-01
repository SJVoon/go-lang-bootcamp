package main

import "fmt"

func main() {

	//switch case without variable and understand fallthrough
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("4 not equal 2")
	case (2 == 2):
		fmt.Println("2 equal 2")
		fallthrough
	case (4 == 4):
		fmt.Println("4 equal 4")
	default:
		fmt.Println("default")
	}

	//switch with a variable
	x := 5
	switch x {
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("default")
	}

}
