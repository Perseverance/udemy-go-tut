package main

import "fmt"

func main() {
	var i int
	fmt.Println("I see you want a FizBuzz. What is your limit human?")
	fmt.Scanln(&i)
	fizbuzz(i)
}

func fizbuzz(limit int) {
	if limit <= 0 {
		fmt.Println("Too short of a number eh?")
		return
	}

	for i := 1; i <= limit; i++ {
		if (i % 15) == 0 {
			fmt.Println("FizBuzz")
		} else if (i % 3) == 0 {
			fmt.Println("Fiz")
		} else if (i % 5) == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
