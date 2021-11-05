package main

import "fmt"

func main() {
	count := 23
	var result string

	if count < 10 {
		result = "regular user"
	} else if count > 10 {
		result = "watch out"
	} else {
		result = "exact"
	}

	fmt.Println(result)

	// initializing and checking together
	if num := 10; num < 10 {
		fmt.Println("less")
	} else if num > 10 {
		fmt.Println("greater")
	} else {
		fmt.Println("equal")
	}
}
