package main

import (
	"fmt"
	"math"
)

func formula(x float64) float64 {
	var a float64 = -2.5
	var b float64 = 3.4
	var i float64 = 0
	if x > 5 {
		var y float64 = (3 * math.Log10(a*a+x)) / ((a + x) * (a + x))
		i = y
	}
	if x < 6 {
		var y float64 = math.Pow((a+b*x), 2.5) / (1.8 + math.Pow(math.Cos(a*x), 3))
		i = y
	}
	return i
}

func main() {
	fmt.Println("Задание А")
	var x float64 = 3.5
	for ; x < 6.5; x = x + 0.6 {
		fmt.Println(formula(x))
	}
	fmt.Println("Задание В")
	fmt.Println(2.89)
	fmt.Println(3.54)
	fmt.Println(5.21)
	fmt.Println(6.28)
	fmt.Println(3.47)
}
