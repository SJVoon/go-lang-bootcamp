package main

import "fmt"

func main() {

	//define by map[key_type]value_type
	m := map[string]int{
		"James": 30,
		"Bond":  32,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barbara"])

	v, ok := m["Barbara"]
	fmt.Println(v)
	fmt.Println(ok)

	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
		// fmt.Println(v)
	}

	delete(m, "James")
	fmt.Println(m)

}
