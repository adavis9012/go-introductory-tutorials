package main

import "fmt"

type float float32

func main() {
	var f float = 52.2
	// var g float32 = 52.2

	fmt.Printf("f has value %v and type %T\n", f, f)

	// This trow an error (mismatched types float and float32)
	// fmt.Println("f == g", f == g)
}
