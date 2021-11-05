package main

import "fmt"

func main() {
	// no inheritance in golang
	subhra := User{"subhra", "subhra@gmail.com", true, 21}
	fmt.Printf("%+v\n", subhra)

	fmt.Println("name: ", subhra.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
