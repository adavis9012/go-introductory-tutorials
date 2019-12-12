package main

import "fmt"

func main() {
	condition := true

	if condition {
		fmt.Println("condition meet")
	}

	fmt.Println("\n----if-else-----")

	a := 2

	if a>10 {
		fmt.Println("condition a>10 meet")
	} else {
		fmt.Println("condition a>10 not  meet")
	}

	fmt.Println("\n----if-else if-else----")
	fruit := "orange"

	if fruit == "mango" {
		fmt.Println("fruit is mango")
	} else if fruit == "orange" {
		fmt.Println("fruit is orange")
	} else if fruit == "banana" {
		fmt.Println("fruit is banana")
	} else {
		fmt.Println("I don't know which fruit this is")
	}


	fmt.Println("\n----initial statement----")

	if fruit2 := "lemon"; fruit2 == "mango" {
		fmt.Println("fruit is mango")
	} else if fruit2 == "orange" {
		fmt.Println("fruit is orange")
	} else if fruit2 == "banana" {
		fmt.Println("fruit is banana")
	} else {
		fmt.Println("I don't know which fruit this is")
	}


	fmt.Println("\n----switch----")
	finger := 2

	switch finger {
	case 1:
		fmt.Println("Thumb");
	case 2:
		fmt.Println("Index");
	case 3:
		fmt.Println("Middle");
	case 4:
		fmt.Println("Ring");
	case 5:
		fmt.Println("Pinky");
	default:
		fmt.Println("No fingers matched")
	}


	fmt.Println("\n----multiple cases values----")
	letter := "i"

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("This letter is a vovel.")
	default:
		fmt.Println("This letter i not a vovel.")
	}



	fmt.Println("\n----initial statement----")

	switch letter2 := "x"; letter2 {
	case "a", "e", "i", "o", "u":
		fmt.Println("This letter is a vovel.")
	default:
		fmt.Println("This letter i not a vovel.")
	}



	fmt.Println("\n----expressionless switch----")

	switch number := 20; {
	case number <= 5:
		fmt.Println("number is less than or equal 5")
	case number > 5:
		fmt.Println("number is greater than 5")
	case number > 10:
		fmt.Println("number is greater than 10")
	case number > 15:
		fmt.Println("number is greater than 15")
	}



	fmt.Println("\n----fallthrough statement----")

	switch number2 := 20; {
	case number2 <= 5:
		fmt.Println("number is less than or equal 5")
		fallthrough
	case number2 > 5:
		fmt.Println("number is greater than 5")
		fallthrough
	case number2 > 10:
		fmt.Println("number is greater than 10")
		fallthrough
	case number2 > 15:
		fmt.Println("number is greater than 15")
	}
}
