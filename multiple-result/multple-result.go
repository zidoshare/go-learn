package main

import (
	"fmt"
)

//swap 可以返回任意数量的值
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
