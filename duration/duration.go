package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	def, _ := time.ParseDuration("1s")
	d := flag.Duration("d", def, "duration time")
	flag.Parse()

	fmt.Println(d.Seconds() + 10) // d+10
}
