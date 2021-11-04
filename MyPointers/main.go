package main

import "fmt"

func main() {
	var ptr *int

	var a int = 2
	ptr = &a

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println(a)
}
