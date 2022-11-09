package main

import (
	"fmt"
	"math"
)

// variant 12

func CalculateEquation(x float64, a float64) float64 {
	var temp float64 = math.Pow(x, 2) - 1
	return math.Pow(a, temp) - math.Log10(temp) + math.Pow(temp, float64(1)/3)
}

func main() {
	var a float64 = 1.6

	// Task A
	fmt.Println("Task A")
	for i := 1.2; i <= 3.7; i += 0.5 {
		fmt.Println("x = ", i, "\ty = ", CalculateEquation(i, a))
	}

	// Task B
	fmt.Println("Task B")
	var x [5]float64 = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	for i := 0; i < 5; i++ {
		fmt.Println("x = ", x[i], "\ty = ", CalculateEquation(x[i], a))
	}
}
