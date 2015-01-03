package math

import (
	"math"
)

func Clamp(value, min, max float64) float64 {
	if value > max {
		return max
	} else if value < min {
		return min
	}
	return value
}
