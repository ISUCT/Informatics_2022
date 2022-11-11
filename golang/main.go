package main

import (
	"fmt"
	"math"
)

func main() {
	println("Задание A")
	var i float64 = 0.26
	for ; i < 0.66; i = i + 0.08 {
		fmt.Println(formula(i))
	}
	fmt.Println("Задание B")
	fmt.Println(formula(0.1))
	fmt.Println(formula(0.3))
	fmt.Println(formula(0.4))
	fmt.Println(formula(0.55))
	fmt.Println(formula(0.6))
}

func formula(x float64) float64 {
	var y float64 = (math.Pow(math.Pow(math.Asin(x), 2)-(math.Pow(math.Acos(x), 3)), 4))
	return y
}
