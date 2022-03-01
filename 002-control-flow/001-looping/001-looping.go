package main

import "fmt"

func main() {

	// for init; condition; post {}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//nested loop
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, j)
		}
	}

	//for single condition, like while loop
	x := 0
	for x < 10 {
		x++
		fmt.Println(x)
	}

	//for without condition, put condition in block
	for {
		if x > 15 {
			break
		}
		x++
		fmt.Println(x)
	}

	//using break and continue
	//break if more than 10 and continue if odd
	y := 0
	for {
		y++
		if y > 10 {
			break
		}
		if y%2 != 0 {
			continue
		}
		fmt.Println(y)
	}

}
