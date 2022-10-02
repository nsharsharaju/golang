package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to class on Mutex and Wait Groups")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	score := []int{0}

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
