package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {

	studentGroup := people{"Zeno", "John", "Al", "Jehny"}

	strings := []string{"Zeno1", "John1", "Al1", "Jehny1"}

	ints := []int{1, 432, 44, 14, 984, -11, 23, 942, 1365, 8462, 73, 6}

	strSliceGroup := sort.StringSlice(studentGroup)

	strSliceGroup.Sort()

	fmt.Println(studentGroup)

	sort.Stable(sort.Reverse(strSliceGroup))

	fmt.Println(studentGroup)

	sort.StringSlice(strings).Sort()

	fmt.Println(strings)

	sort.IntSlice(ints).Sort()

	fmt.Println(ints)
}
