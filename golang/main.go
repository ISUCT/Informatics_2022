package main

import (
	"fmt"
	"math"
)

func CalculateEquation(a, b, x float64) float64 {
	return math.Sqrt(math.Abs(a-b*x) / math.Pow(math.Log10(x), 3))
}

func main() {
	fmt.Println("TASK A")
	a := 7.2
	b := 4.2
	for i := 1.81; i <= 5.31; i += 0.7 {
		fmt.Println("x = ", i, "\ty = ", CalculateEquation(i, a, b))
	}

	fmt.Println("TASK B")
	x := [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	for i := 0; i < 5; i++ {
		fmt.Println("x = ", x[i], "\ty = ", CalculateEquation(a, b, x[i]))
	}
}
