package main

import "fmt"

func main() {
	fmt.Print("You want the sum of numbers divisible to 3 or 5 up to : ")
	var limit int

	fmt.Scan(&limit)

	fmt.Println(sum3to5(limit))
}

func sum3to5(limit int) int {
	if limit <= 0 {
		fmt.Println("Too short of a number eh?")
		return 0
	}

	var sum int

	for i := 1; i < limit; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return sum
}
