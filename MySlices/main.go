package main

import (
	"fmt"
)

func main() {
	// var l = []string{"apple", "mango"}
	// fmt.Printf("%T\n", l)
	// l = append(l, "banana", "nuts")

	// fmt.Println(l[:2])

	// scores := make([]int, 4)

	// scores[0] = 323
	// scores[1] = 434
	// scores[2] = 898
	// scores[3] = 454

	// scores = append(scores, 676, 897, 264)
	// fmt.Println(scores)

	// sort.Ints(scores)
	// fmt.Println(scores)

	// removing value with index from slice
	var cources = []string{"js", "go", "java", "c++"}
	fmt.Println(cources)

	var i int = 2
	cources = append(cources[:i], cources[i+1:]...)
	fmt.Println(cources)
}
