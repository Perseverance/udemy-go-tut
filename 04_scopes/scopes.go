package main

import (
	"fmt"
)

var x = 42

const (
	cPi = 3.14
)

func main() {
	adder := returnFunc(x)
	fmt.Println(adder(11))
	fmt.Println(adder(13))
	fmt.Println(cPi)
}

func returnFunc(s int) func(int) int {
	retFn := func(x int) int {
		return s + x
	}
	return retFn
}
