package main

import (
	"fmt"
	"math"
)

func main() {
	const a float64 = 2.0
	const b float64 = 1.1
	task1(a, b)
	fmt.Println(task2(a, b))
}

func formula(x float64, a float64, b float64) float64 {
	var y float64 = (math.Log10(math.Abs(b*b-x*x)) / math.Log10(a)) / math.Pow(math.Abs(x*x-a*a), 0.2)
	return y
}

func task1(a float64, b float64) {
	fmt.Println("Задание А")
	var x float64 = 0.08
	for ; x < 1.08; x = x + 0.2 {
		fmt.Println(formula(x, a, b))
	}
}

func task2(a float64, b float64) [5]float64 {
	fmt.Println("Задание В")
	var Massiv = [5]float64{formula(0.1, a, b), formula(0.3, a, b), formula(0.4, a, b), formula(0.45, a, b), formula(0.65, a, b)}
	return Massiv
}
