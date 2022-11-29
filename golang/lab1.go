package main

import (
	"math"
)

func function(a, x float64) float64 {
	var y float64 = math.Pow(a, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Pow((math.Pow(x, 2)-1), (1.0/3.0))
	return y
}

func zadacha1(a, xn, xk, dx float64) []float64 {
	var zadacha []float64
	for x := xn; x < xk; x += dx {
		zadacha = append(zadacha, function(a, x))
	}
	return zadacha
}

func zadacha2(a float64, znach []float64) []float64 {
	var zadachaB []float64
	for _, x1 := range znach {
		zadachaB = append(zadachaB, function(a, x1))
	}
	return zadachaB
}
