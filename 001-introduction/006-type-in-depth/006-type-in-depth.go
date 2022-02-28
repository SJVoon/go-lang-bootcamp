package main

import "fmt"

func main() {
	boolFunc()
	numericFunc()
	stringFunc()
	constFunc()
	iotaFunc()
}

var x bool

func boolFunc() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 40
	b := 42
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}

var a int
var b float64

func numericFunc() {
	a = 123
	b = 123.232
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// byte alias of uint8
	// rune alias of int32
}

func stringFunc() {
	//raw string literal
	str := `"all string here format etc is fixed"`
	s := "Hello, Programming"
	bs := []byte(s)
	fmt.Println(str)
	fmt.Println(s)
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}
}

const c1 = 42
const c2 float64 = 42.78
const (
	c3        = 65
	c4 string = "hello"
)

func constFunc() {
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
}

const (
	_d = iota
	d1 = iota
	d2
	d3
)

//bit shifting using iota
//y = x << 1
//y is equal to x shifting 1 bit to left side
//see example below, b1 = 1(2^0) shift 1 bit => 2(2^1)
//see example below, b1 = 1(2^0) shift 1 bit => 4(2^2)
//<< is left shift, >> is right shift
const (
	b0 = iota
	b1 = 1 << iota
	b2 = 1 << iota
	b3 = b2 >> 1
)

func iotaFunc() {
	fmt.Println(_d)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Printf("%d\t\t%b\n", b0, b0)
	fmt.Printf("%d\t\t%b\n", b1, b1)
	fmt.Printf("%d\t\t%b\n", b2, b2)
	fmt.Printf("%d\t\t%b\n", b3, b3)
}
