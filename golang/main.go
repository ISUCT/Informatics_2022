package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		A float64 = 2.5
		B float64 = 4.6
	)
	fmt.Println("Вывод задачи А по формуле 18:")
	task_a(A, B, 1.1, 3.6, 0.5)
	fmt.Println("Вывод задачи B по формуле 18:")
	variables := [5]float64{1.2, 1.28, 1.36, 1.46, 2.35}
	task_b(A, B, variables[:])
}

func task_a(a, b, xn, xk, dx float64) {
	for x := xn; x <= xk; x += dx {
		answer := formula18(a, b, x)
		fmt.Println(answer)
	}
}

func task_b(a, b float64, xs []float64) {
	for _, x := range xs {
		answer := formula18(a, b, x)
		fmt.Println(answer)
	}
}

func formula18(a, b, x float64) float64 {
	answer := math.Pow((a+b*x), 2.5) / (1 + math.Log10(a+b*x))
	return answer
}
