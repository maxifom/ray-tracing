package utils

import (
	"math"
	"math/rand"
)

func Schlick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func RandomDouble(a, b float64) float64 {
	return a + (b-a)*rand.Float64()
}
