package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("___Задание А___")
	var x float64 = 1.2
	for ; x < 4.2; x = x + 0.6 {
		fmt.Println(formula(x))
	}
	fmt.Println("___Задание В___")
	fmt.Println(formula(1.16))
	fmt.Println(formula(1.32))
	fmt.Println(formula(1.47))
	fmt.Println(formula(1.65))
	fmt.Println(formula(1.93))
}

func formula(x float64) float64 {
	var a float64 = 2.0

	var y float64 = (math.Log10(a + x)) / ((a + x) * (a + x))
	return y
}
