package main

import (
	"fmt"
	"math"
)

func formula(a, b, x float64) float64 {
	y := (math.Pow((a*x + b), 1/3)) / (math.Pow(math.Log10(x), 2))
	return y
}

func TaskA(a, b, xn, xk, dx float64) []float64 {
	var out []float64
	for x := xn; x <= xk; x += dx {
		out = append(out, formula(a, b, x))
	}
	return out
}

func TaskB(a, b float64, massiv []float64) []float64 {
	var out []float64
	for _, x1 := range massiv {
		out = append(out, formula(a, b, x1))
	}
	return out
}

func main() {
	var a = 1.35
	var b = 0.98
	var xn = 1.14
	var xk = 4.24
	var dx = 0.62
	fmt.Println("Задание a")
	fmt.Println("y = ", TaskA(a, b, xn, xk, dx))
	fmt.Println("Задание b")
	var massiv = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
	fmt.Println("y = ", TaskB(a, b, massiv[:]))
}
