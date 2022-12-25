package internal

import (
	"fmt"
	"math"
)

func TestAB(a float64, b float64) {
	Task1(0.08, 0.2, 1.08, a, b)
	Task2([]float64{0.1, 0.3, 0.4, 0.45, 0.65}, a, b)
}

// Формула для Заданий А и Б
func formula(x float64, a float64, b float64) float64 {
	var y float64 = (math.Log10(math.Abs(b*b-x*x)) / math.Log10(a)) / math.Pow(math.Abs(x*x-a*a), 0.2)
	return y
}

// Задание А
func Task1(xn float64, xd float64, xk float64, a float64, b float64) {
	fmt.Println("Задание А")
	var x float64 = xn
	for ; x < xk; x = x + xd {
		fmt.Println(formula(x, a, b))
	}
}

// Задание Б
func Task2(Massiv []float64, a float64, b float64) {
	fmt.Println("Задание В")
	for i := 0; i < 5; i++ {
		Massiv[i] = formula(Massiv[i], a, b)
	}
	fmt.Println(Massiv)
}
