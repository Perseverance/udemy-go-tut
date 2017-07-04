package main

import (
	"fmt"
)

func main() {
	var a int
	var b rune
	var c string
	var d float64
	var e bool

	fmt.Printf("%T : %v\n", a, a)
	fmt.Printf("%T : %v\n", b, b)
	fmt.Printf("%T : %v\n", c, c)
	fmt.Printf("%T : %v\n", d, d)
	fmt.Printf("%T : %v\n", e, e)
}
