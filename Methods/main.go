package main

import "fmt"

func main() {
	// no inheritance in golang
	subhra := User{"subhra", "subhra@gmail.com", true, 21}
	fmt.Printf("%+v\n", subhra)

	subhra.GetStatus()

	subhra.SetNewEmail()
	fmt.Println(subhra.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Status:", u.Status)
}

func (u User) SetNewEmail() {
	u.Email = "new@gmail.com"
	fmt.Println(u.Email)
}
