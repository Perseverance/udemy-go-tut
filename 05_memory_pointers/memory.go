package main

import "fmt"

func main() {
	u := 42

	i := &u

	increment(&u)
	fmt.Printf("%T, %v %v\n", i, i, *i)

	increment(i)
	fmt.Printf("%T, %v %v\n", i, i, *i)

	fmt.Printf("%T, %v \n", u, u)

}

func increment(x *int) {
	fmt.Printf("%T, %v %v\n", x, x, *x)
	*x++
}
