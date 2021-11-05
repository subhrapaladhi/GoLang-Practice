package main

import "fmt"

func main() {
	greeter()

	result := adder(3, 6)
	fmt.Println(result)

	result2 := addMultiple(3, 6, 5, 7, 1, 6, 2, 2, 75, 3, 23, 34, 23)
	fmt.Println(result2)

	name, age := getNameAge()
	fmt.Println(name, age)
}

func greeter() {
	fmt.Println("Hello world")
}

func adder(a int, b int) int {
	return a + b
}

func addMultiple(vals ...int) int {
	total := 0

	for _, val := range vals {
		total += val
	}
	return total
}

func getNameAge() (string, int) {
	return "subhra", 21
}
