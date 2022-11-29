package main

import (
	"math"
)

func function(a, x float64) float64 {
	var y float64 = math.Pow(a, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Pow((math.Pow(x, 2)-1), (1.0/3.0))
	return y
}

func zadacha1(a float64) []float64 {
	var zadacha []float64
	for x := 1.2; x < 4.2; x = x + 0.6 {
		zadacha = append(zadacha, function(a, x))
	}
	return zadacha
}

func zadacha2(a float64) []float64 {
	var zadachaB []float64
	var znach = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	for _, x1 := range znach {
		zadachaB = append(zadachaB, function(a, x1))
	}
	return zadachaB
}
