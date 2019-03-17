package main

import (
	"log"
)

func main() {
	log.SetFlags(0)
	log.Println("this is a log") // this is a log

	log.SetFlags(log.LstdFlags)
	log.Println("this is a log") // 2018/12/31 01:02:54 this is a log

	log.SetFlags(log.Ldate)
	log.Println("this is a log") // 2018/12/31 this is a log

	log.SetFlags(log.Ltime)
	log.Println("this is a log") // 01:02:54 this is a log

	log.SetFlags(log.Lmicroseconds)
	log.Println("this is a log") // 01:04:37.350084 this is a log

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("this is a log") // 2018/12/31 01:02:54.289191 log.go:21: this is a log
}
