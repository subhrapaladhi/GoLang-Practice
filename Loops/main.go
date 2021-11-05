package main

import "fmt"

func main() {
	days := []string{"sun", "mon", "tue", "wed", "thur", "fri", "sat"}

	fmt.Println(days)

	// for loop
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// like foreach loop
	// for idx, day := range days {
	// 	fmt.Println(idx, ": ", day)
	// }

	// like while loop
	// r := 0
	// for r < 10 {
	// 	fmt.Println(r)
	// 	r++

	// 	if r == 5 {
	// 		break
	// 	}
	// }

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++

		if i == 5 {
			goto p1
		}
	}

p1:
	fmt.Println("goto statement p1")
}
