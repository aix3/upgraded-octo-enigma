package main

import (
	"log"
	"sync"
)

func main() {
	ch := make(chan int, 1)

	var wg sync.WaitGroup

	var fn func(ch chan int)
	fn = func(ch chan int) {
		wg.Add(1)
		// defer wg.Done()
		val := <-ch
		val++
		log.Println("val++:", val)
		ch <- val
		go fn(ch)
	}

	ch <- 1
	fn(ch)

	wg.Wait()
}
