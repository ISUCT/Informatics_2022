package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задание А")
	var i float64 = 0.08
	for ; i < 1.08; i = i + 0.2 {
		fmt.Println(formula(i))
	}
	fmt.Println("Задание В")
	fmt.Println(formula(0.1))
	fmt.Println(formula(0.3))
	fmt.Println(formula(0.4))
	fmt.Println(formula(0.45))
	fmt.Println(formula(0.65))
}

func formula(x float64) float64 {
	var a float64 = 2.0
	var b float64 = 1.1
	var y float64 = (math.Log10(math.Abs(b*b-x*x)) / math.Log10(a)) / math.Pow(math.Abs(x*x-a*a), 0.2)
	return y
}
