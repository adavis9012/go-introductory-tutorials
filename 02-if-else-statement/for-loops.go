package main

import "fmt"

func main() {
	for i :=1; i <= 2; i++ {
		fmt.Printf("current number is %d\n", i)
	}


	fmt.Println("\n\nVARIANTS")
	fmt.Println("\n--optional post statement--")
	for i := 1; i <= 2; {
		fmt.Printf("current number is %d\n", i)
		i++
	}

	fmt.Println("\n--optional init statement--")
	j := 1
	for ; j <= 2; j++ {
		fmt.Printf("current number is %d\n", j)
	}

	fmt.Println("\n--optional init and post statement--")
	k := 1
	for k <= 2 {
		fmt.Printf("current number is %d\n", k)
		k++
	}

	fmt.Println("\n--without statements, break statement--")
	l := 1
	for {
		fmt.Printf("current number is %d\n", l)
		if l == 2 {
			break;
		}
		l++
	}

	fmt.Println("\n--continue statement--")
	for i := 1; i <= 4; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("current number is %d\n", i)
	}
}
