package main

import "fmt"

type person struct {
	first string
	last  string
}

//embedded struct person in here
type secretAgent struct {
	person
	code string
}

func main() {

	//create a value of type person
	//similar meaning to create an object of a class
	p1 := person{
		first: "first",
		last:  "people",
	}

	p2 := person{
		first: "james",
		last:  "bond",
	}

	sa1 := secretAgent{
		person: p2,
		code:   "007",
	}

	//anonymous struct declare inline
	anonymous1 := struct {
		first string
		last  string
		age   int
	}{
		first: "anony",
		last:  "mous",
		age:   32,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa1)
	fmt.Println(anonymous1)

	fmt.Println(p1.first, p1.last, p2.first, p2.last)
}
