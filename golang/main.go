package main

import (
	"fmt"
	"math"
)

func main() {
	const a float64 = 2.5
	const b float64 = 4.6
	fmt.Print("Задача A\n")
	for i := 1.1; i < 3.6; i += 0.5 {
		var result float64 = task(a, b, i)
		fmt.Printf("Результат вычислений = %.4f\n", result)
	}
	fmt.Print("\n")
	fmt.Print("Задача B\n")
	array := [5]float64{1.2, 1.28, 1.36, 1.46, 2.35}
	for _, element := range array {
		var result float64 = task(a, b, element)
		fmt.Printf("Результат вычислений = %.4f\n", result)
	}
}

func task(a, b, x float64) float64 {
	return math.Pow((a+b*x), 2.5) / (1 + math.Log(a+b*x))
}
