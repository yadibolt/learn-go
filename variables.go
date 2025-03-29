package main

import "fmt"

func main() {
	// without initialization
	var foo string
	// with initialization
	var foo string = "Go String Test"
	// multiple - line declarations
	var foo, bar string = "Foo", "Bar"
	var (
		foo string = "Hello World"
		bar string = "World Hello"
	)

	// data types
	var a string = "a"
	var b bool = false
	
	// int dependant on the system
	var i int = 200
	// int and uint to 64 (8, 16, 32, 64)

	// byte and rune
	var b byte = 'a'
	var r rune = '👏'

	// f32 and f64...
	// ...
	var i int
	var f float64
	var b bool
	var s string
	
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
