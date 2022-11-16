package main

import "math"

func lab(a float64, b float64, x float64) float64 {
	v2 := math.Pow(a+b*x, 2.5) / (1.8 + math.Pow(math.Cos(a*x), 3))
	if (x > 5) && (v2 >= 5) {
		y := math.Pow(math.Log10(math.Pow(a, 2)-x), 2) / math.Pow(a+x, 2)
		return y
	} else {
		return 0
	}
}
