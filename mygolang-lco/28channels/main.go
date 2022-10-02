package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to class on Channels")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("In 1")
		fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	wg.Add(1)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("In 2")
		ch <- 2
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
