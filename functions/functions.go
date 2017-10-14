package main

import "fmt"

//add 注意类型在变量名 _之后_
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(2, 5))
}
