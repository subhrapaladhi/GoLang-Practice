package main

import "fmt"

func main() {
	var l [4]int

	l[0] = 23
	l[1] = 77
	l[2] = 34

	fmt.Println(l)

	var l1 = [3]int{25, 56}
	fmt.Println(l1)
}
