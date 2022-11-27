package main

import (
	"math"
)

func function(a, x float64) float64 {
	var y float64 = math.Pow(a, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Pow((math.Pow(x, 2)-1), (1.0/3.0))
	return y
}

func zadacha1(a, x float64) float64 {
	var f float64 = function(a, x)
	return f
}

func zadacha2(a, x1 float64) float64 {
	var f1 float64 = function(a, x1)
	return f1
}
