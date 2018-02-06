package main

import (
	"io"
	"strings"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(bs []byte) (int, error) {
	i, err := rot.r.Read(bs)
	rot13(bs)
	return i, err
}

func rot13(bs []byte) {
	for i := range bs {
		bs[i] = get13(bs[i])
	}
}

func get13(b byte) byte {
	if b >= 'a' && b <= 'z' {
		r := 'z' - b
		if r >= 13 {
			return b + 13
		} else {
			return b - 13
		}
	}
	if b >= 'A' && b <= 'Z' {
		r := 'Z' - b
		if r >= 13 {
			return b + 13
		} else {
			return b - 13
		}
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr! n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
