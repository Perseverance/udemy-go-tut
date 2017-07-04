package main

import (
	"fmt"

	"github.com/Perseverance/udemy-go-tut/02_packages/stringutil"
)

var test = "asd123"

func main() {
	fmt.Println(stringutil.Reverse(test))
}
