package main

import "math"

func lab(a float64, x float64) float64 {
	y := math.Pow(math.Log10(a+x), 2) / math.Pow((a+x), 2)
	return y
}
