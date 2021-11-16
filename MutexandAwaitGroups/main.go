package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	var score = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("one")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)
	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("two")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("three")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(score)
}
