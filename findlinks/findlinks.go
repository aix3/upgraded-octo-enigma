package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%v\n", string(b))
}
