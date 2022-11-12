package main

import (
	"fmt"
	"math"
)

func formula(x float64) float64 {
	var a float64 = 4.1
	var b float64 = 2.7
	var y float64 = (a*(math.Cbrt(x)) - (b * (math.Log(x) / math.Log(5)))) / (math.Log10(math.Abs(x - 1)))
	return y
}

func main() {
	fmt.Println("Задание А")
	var i float64 = 1.2
	for ; i < 5.2; i = i + 0.8 {
		fmt.Println(formula(i))
	}
	fmt.Println("Задание B")
	fmt.Println(formula(1.9))
	fmt.Println(formula(2.15))
	fmt.Println(formula(2.34))
	fmt.Println(formula(2.73))
	fmt.Println(formula(3.16))
}
