package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type sphere struct {
	side float64
}

func (s sphere) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	s := sphere{
		side: 12,
	}

	c := circle{
		radius: 4,
	}

	info(s)

	info(c)

}
