package sprint

import "math"

func Casting(n float64) int {
	x := math.Round(n)
	return int(x)
}