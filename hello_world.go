package main

import "fmt"

// Declare a variable outside of main()
var a int

// declare group of variables
var (
	b bool
	c float32
	d string
)

func main() {
	fmt.Println("Hello, Paraguay!")

	a = 42
	b, c = true, 29.5
	d = "this is a string, dog" // must be double quotes

	fmt.Println(a, b, c, d)

	// now with short declarations babyy
	f := 23
	g, h := true, 55.5
	i := "Anotha one"

	fmt.Println(f, g, h, i)

}
