package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("hello %g problems.", math.Nextafter(2, 3))
}
