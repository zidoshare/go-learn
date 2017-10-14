package main

import (
	"fmt"
)

//fplit 返回值可以被命名，并且像变量那样使用。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(split(17))
}
