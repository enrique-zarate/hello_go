package main

import "fmt"

func main() {
	// especificado por el usuario
	const a int32 = 23
	const b float32 = 9.99
	var c complex64 = 1 + 4i
	var d uint16 = 16 // 16-bit unsigned integer

	// default types
	n := 44
	pi := 3.14
	x, y := true, false
	z := "Aguante Gooooo"

	fmt.Printf("especificados por el usuario:\n%T, %T,%T,%T", a, b, c, d)

	fmt.Printf("\n\ndefault: \n%T, %T, %T, %T,%T ", n, pi, x, y, z)
}
