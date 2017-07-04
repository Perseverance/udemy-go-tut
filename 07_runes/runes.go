package main

import "fmt"

func main() {
	printStrings()
}

func printHello() {
	hello := "Hello World			"
	fmt.Printf("%v \n", []int32(hello))
}

func printChars() {
	for i := 11040; i < 11140; i++ {
		x := string(i)
		fmt.Printf("%v - %v - %v\n", i, x, []byte(x))
	}
}

func printStrings() {
	fmt.Println(
		`Todd
Is
A
Beach`)
}
