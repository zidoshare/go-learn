package main

import (
	"math"
	"golang.org/x/tour/pic"
)
//Pic (i+j)/2
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8,dy)
	for i := range result {
		result[i] = make([]uint8,dx)
		for j := range result[i] {
			result[i][j] = uint8(math.Pow(float64(i),float64(j)))
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}