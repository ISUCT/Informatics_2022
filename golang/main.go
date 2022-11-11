package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1.1
	var b float64 = 0.09

	fmt.Printf("Задача А\n")
	for i := 1.2; i < 2.2; i += 0.2 {
		Problem(a, b, i)
	}
	fmt.Printf("-------------------------------\n")

	fmt.Printf("Задача B\n")

	xValues := [5]float64{1.21, 1.76, 2.53, 3.48, 4.52}
	for _, element := range xValues {
		Problem(a, b, element)
	}
}

func Problem(a, b, x float64) {
	var result float64 = math.Log(x*x-1) / Log(5, (a*x*x-b))
	fmt.Printf("a = %.4f, b = %.4f, x = %.4f, result = %.4f\n", a, b, x, result)
}

func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}
