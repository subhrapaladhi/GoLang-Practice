package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))

	createdDate := time.Date(2020, time.November, 3, 23, 23, 23, 23, time.UTC)
	fmt.Println(createdDate)
}
