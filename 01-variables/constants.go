package main

import "fmt"

func main() {
	const(
		a=1
		b=2
		c
		d
	)

	const(
		e= iota
		f= iota
		g= iota
		h
		_
		_
		i
	)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	fmt.Println("/////////iota/////////")

	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
	fmt.Println("i: ", i)
}
