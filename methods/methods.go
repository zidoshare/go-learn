package main

import (
	"fmt"
	"math"
)

// Vertex .
type Vertex struct {
	X, Y float64
}

/*
Abs Go 没有类。然而，仍然可以在结构体类型上定义方法。

方法接收者 出现在 func 关键字和方法名之间的参数中。
*/
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//MyFloat 你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
//但是，不能对来自其他包的类型或基础类型定义方法。
type MyFloat float32

//Abs .
func (f MyFloat) Abs() float32 {
	return float32(f * f)
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(v.Abs())

	var x MyFloat = 10
	y := MyFloat(10)

	fmt.Println(x.Abs(), y)
}
