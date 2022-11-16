package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задание А")
	var i float64 = 1.5
	for ; i < 3.5; i = i + 0.4 {
		fmt.Println(formula(i))
	}
	fmt.Println("Задание В")
	fmt.Println(formula(1.9))
	fmt.Println(formula(2.15))
	fmt.Println(formula(2.34))
	fmt.Println(formula(2.74))
	fmt.Println(formula(3.16))
}

func formula(x float64) float64 {
	var a float64 = 4.1
	var b float64 = 2.7
	var y float64 = (a*(math.Cbrt(x)) - b*(math.Log10(x)/math.Log10(5))) / (math.Pow(math.Log10(x-1), 3))
	return y
}
