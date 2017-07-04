package main

import "fmt"

func main() {
	var name string
	fmt.Println("Hey! What is your name?")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s \n", name)
}
