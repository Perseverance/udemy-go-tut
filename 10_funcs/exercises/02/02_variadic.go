package main

import "fmt"

func main() {
	fmt.Println(findGreatest(11, 44, 872378, 563126, 72910, 784312, 871234))
	fmt.Println(findGreatest(1, 2))
	fmt.Println(findGreatest(1, 2, 3))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(findGreatest(aSlice...))
	fmt.Println(findGreatest())
}

func findGreatest(nums ...int) int {
	var greatest int
	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}
