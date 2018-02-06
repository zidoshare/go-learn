package main

import (
	"fmt"
	"math"
)

type MFloat float64

func (f MFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
