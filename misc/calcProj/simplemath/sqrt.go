package simplemath

import "math"

func Sqrt(i int) float32 {
	v := math.Sqrt(float64(i))
	return float32(v)
}
