package helpers

import "math"

func RoundUp(value float32) float32 {
	return float32(math.Ceil(float64(value)*10000) / 10000)
}
