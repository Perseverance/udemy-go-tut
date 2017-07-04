package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d%d\n", i, j)
		}
	}
}
