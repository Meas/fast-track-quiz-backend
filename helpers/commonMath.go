package helpers

import "math"

// RoundUp - rounds up a float32 to 4 decimal places
func RoundUp(value float32) float32 {
	return float32(math.Ceil(float64(value)*10000) / 10000)
}
