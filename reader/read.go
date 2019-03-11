package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.pptbz.com/pptpic/UploadFiles_6909/201203/2012031220134655.jpg")
	b := make([]byte, resp.ContentLength)
	n, err := resp.Body.Read(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ContentLength, n, b[0:10], b[len(b)-10:len(b)]) // 64093 3720 [255 216 255 224 0 16 74 70 73 70] [0 0 0 0 0 0 0 0 0 0]

	resp, _ = http.Get("http://www.pptbz.com/pptpic/UploadFiles_6909/201203/2012031220134655.jpg")
	lr := io.LimitReader(resp.Body, resp.ContentLength)
	b, _ = ioutil.ReadAll(lr)

	fmt.Println(resp.ContentLength, len(b), b[0:10], b[len(b)-10:len(b)]) // 64093 64093 [255 216 255 224 0 16 74 70 73 70] [208 127 115 57 103 56 247 63 255 217] WHY?
}
