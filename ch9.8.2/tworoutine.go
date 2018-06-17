package main

import "time"
import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	var times int64
	trans := func(src, dst chan string) {
		for {
			msg := <-src
			times++
			dst <- msg
		}
	}

	go func() {
		ch1 <- "Hello Go!"
	}()

	go trans(ch1, ch2)
	go trans(ch2, ch1)

	time.Sleep(time.Second * 1)
	fmt.Println("times:", times)
}
