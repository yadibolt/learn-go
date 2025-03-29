package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	x := 42
	y := float64(x)
	u := uint(y)

	fmt.Printf("%T %T", y, u)
}
