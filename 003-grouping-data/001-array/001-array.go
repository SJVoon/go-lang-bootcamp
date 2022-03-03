package main

import "fmt"

func main() {
	var y [5]int
	fmt.Println(y)
	fmt.Println(len(y))

	x := []int{4, 5, 6, 7, 8}
	for i, v := range x {
		fmt.Println(i, v)
	}

	//using make
	arr := make([]string, 50, 50)
	fmt.Println(arr)

	//2d array or slice of slice
	x1 := []int{4, 5, 6, 7, 8}
	x2 := []int{9, 10, 11, 12, 12}
	fmt.Println(x1)
	fmt.Println(x2)

	xxs := [][]int{x1, x2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, v := range xs {
			fmt.Printf("\t index position: %v\t value: \t %v \n", j, v)
		}
	}
}
