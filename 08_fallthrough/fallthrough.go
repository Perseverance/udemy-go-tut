package main

import "fmt"

func main() {
	switch "John" {
	case "Mike":
		fmt.Println("Hi Mike")
	case "John":
		fallthrough
	case "Bro":
		fmt.Println("Whats up bro")
	}
}
