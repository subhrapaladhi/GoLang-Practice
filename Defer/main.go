package main

import "fmt"

func main() {
	defer fmt.Println("hello world")
	fmt.Println("normal print")
	myDefer()
	fmt.Println(temp())
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func temp() int {
	if 2 < 3 {
		return 100
	} else {
		return 34
	}
}
