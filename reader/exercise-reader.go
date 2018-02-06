package main

import (
	"fmt"
	"io"
)

type MyReader struct {
}

func (r *MyReader) Read(bs []byte) (int, error) {
	l := len(bs)
	for i, _ := range bs {
		bs[i] = 'A'
	}
	return l, nil
}

func main() {
	b := make([]byte, 8)
	r := &MyReader{}
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %c\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
