package main

import (
	"fmt"
	"math"
)

func main() {
	const A, B = 2.5, 4.6
	fmt.Println("Вывод задачи А по формуле 18:")
	task_a(A, B)
	fmt.Println("Вывод задачи B по формуле 18:")
	task_b(A, B)
}

func task_a(A, B float64) {
	for x := 1.1; x <= 3.6; x += 0.5 {
		answer := formula18(A, B, x)
		fmt.Println(answer)
	}
}

func task_b(A, B float64) {
	variables := [5]float64{1.2, 1.28, 1.36, 1.46, 2.35}
	for _, x := range variables {
		answer := formula18(A, B, x)
		fmt.Println(answer)
	}
}

func formula18(a, b, x float64) float64 {
	answer := math.Pow((a+b*x), 2.5) / (1 + math.Log10(a+b*x))
	return answer
}
