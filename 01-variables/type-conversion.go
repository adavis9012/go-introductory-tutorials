package main

import (
	"fmt"
)

func main() {
	var1 := 10
	var2 := 10.5

	// illegal
	// var3 := var1 + var2

	// legal
	var3 := var1 + int(var2)
	var4 := float64(var1) + var2

	fmt.Println("int: ", var1, "+", var2, "=", var3)
	fmt.Println("float64: ", var1, "+", var2, "=", var4)
}
