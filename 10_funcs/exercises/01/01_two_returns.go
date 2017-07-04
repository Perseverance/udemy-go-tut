package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter a num to split in half : ")
	fmt.Scan(&a)

	half := func(num int) (float64, bool) {
		return (float64(num) / 2), (num%2 == 0)
	}

	x, z := half(a)

	fmt.Println(x, z)
}
