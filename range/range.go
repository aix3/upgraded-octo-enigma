package main

import (
	"bytes"
	"fmt"
)

func main() {
	list := []rune{
		'a', 'b', 'c', 'd',
	}

	buf := bytes.NewBuffer([]byte{})
	for i, r := range list {
		fmt.Printf("i=%v, r=%v\n", i, r)
		if i != 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(r)
	}
	fmt.Println(buf.String())
}
