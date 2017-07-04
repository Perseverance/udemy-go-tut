package main

import "fmt"

func main() {
	// fmt.Println(variadic(40, 20, 15, 5))
	// variadicArg := []float64{12, 24, 24, 12}
	// fmt.Println(variadic(variadicArg...))
	// funfun := returningFuncFunc()
	// fun := funfun()
	// fmt.Println(fun())
	withCallback(reusableCallback)
}

func variadic(nums ...float64) float64 {
	var total float64
	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}

func returningFuncFunc() func() func() string {
	return func() func() string {
		return func() string {
			return "Finally"
		}
	}
}

func withCallback(cb func(string)) {
	cb("LOL")
}

func reusableCallback(result string) {
	fmt.Println(result)
}
