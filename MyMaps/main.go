package main

import "fmt"

func main() {
	langs := make(map[string]string)

	langs["js"] = "javascript"
	langs["go"] = "golang"
	langs["py"] = "python"

	fmt.Println(langs)
	fmt.Println(langs["js"])

	delete(langs, "py")
	fmt.Println(langs)

	// looping through maps
	for key, value := range langs {
		fmt.Println(key, ": ", value)
	}
}
