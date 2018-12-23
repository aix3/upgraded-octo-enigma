package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	log.SetPrefix("sss:")
	log.SetFlags(0)
}

func main() {
	seconds := 10
	ten := time.Duration(seconds) * time.Second
	ten1 := time.Second * 10

	fmt.Println(ten, ten1)
	fmt.Println(ten == ten1)
	log.Fatal(waitForServer(os.Args[1]))
}

func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Println("server error:", err, "retrying...")
		sec := time.Second
		log.Println("time.Second =", sec, ", tries =", tries, ", sec << tries =", sec<<uint(tries))
		time.Sleep(sec << uint(tries))
	}
	return fmt.Errorf("server %s faield to respond after %s", url, timeout)
}
