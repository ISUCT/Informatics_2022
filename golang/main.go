package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задание А")

	var xk float64 = 2.2
	var a float64 = 1.2
	var b float64 = 0.48
	for x := 0.7; x < xk; x = x + 0.3 {
		var y float64 = (((math.Pow(b, 3)) + math.Pow(math.Sin(a*x), 2)) / ((math.Acos(x * b * x)) + (math.Pow(2.7183, -x/2))))
		fmt.Println(y)
	}
	fmt.Println("Задание Б")
	var numbersx [5]float64 = [5]float64{0.25, 0.36, 0.56, 0.94, 1.28}
	var c int64 = 0
	for ; c <= 4; c = c + 1 {
		var y float64 = (((math.Pow(b, 3)) + math.Pow(math.Sin(a*numbersx[c]), 2)) / ((math.Acos(numbersx[c] * b * numbersx[c])) + (math.Pow(2.7183, -numbersx[c]/2))))
		fmt.Println(y)

	}
}
