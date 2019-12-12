package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 145.8, 453.8
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)

	// Numeric Expressions
	var na = 11/2
	var nb = 11.0/2
	var nc = float32(11)/2

	fmt.Println("------Numeric Expression-------")
	fmt.Printf("na value %v, type %T\n", na, na)
	fmt.Printf("nb value %v, type %T\n", nb, nb)
	fmt.Printf("nc value %v, type %T\n", nc, nc)
}
