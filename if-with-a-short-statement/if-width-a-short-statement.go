package statement

import (
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); x < lim {
		return v
	}
	return lim
}
