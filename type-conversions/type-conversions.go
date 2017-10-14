package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	//Go 的在不同类型之间的项目赋值时需要显式转换。
	var f = math.Sqrt(float64(x*x + y*y))
	// 表达式 T(v) 将值 v 转换为类型 `T`。
	var z = int(f)

	fmt.Println(x, y, z)
}
