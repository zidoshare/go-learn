package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
// 但是，不能对来自其他包的类型或基础类型定义方法。
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) swap() *Vertex {
	temp := v.X
	v.X = v.Y
	v.Y = temp
	return v
}

func (v *Vertex) equals(other *Vertex) (bool, *Vertex, *Vertex) {
	return other.X == v.X && other.Y == v.Y, v, other
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v)
	fmt.Println(v.swap())
	v.Scale(5)
	fmt.Println(v.Abs())

	other := &Vertex{20, 15}

	result, q, w := v.equals(other)
	fmt.Println(result, q, w)
}
