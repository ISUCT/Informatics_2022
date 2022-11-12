package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задание А")
	var a float64 = 1.35
	var b float64 = 0.98
	var xk float64 = 4.24
	var x float64 = 1.14
	for ; x < xk; x = x + 0.62 {
		var y float64 = (math.Pow(a*x+b, 1.0/3)) / (math.Pow((math.Log10(x)), 2))
		fmt.Println(y)
	}
	fmt.Println("Задание Б")
	var numbersx [5]float64 = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
	var c int64 = 0
	for ; c <= 4; c = c + 1 {
		var y float64 = (math.Pow(a*numbersx[c]+b, 1.0/3)) / (math.Pow((math.Log10(numbersx[c])), 2))
		fmt.Println(y)
	}
}
