//Анастасия Фокина, варинт 21, гр 1/147

package main

import (
	"fmt"
	"math"
)

func equation(x float64) float64 {
	return (math.Pow(math.Sin(x), 3) + math.Pow(math.Cos(x), 3)) * math.Log(x)
}

func main() {

	fmt.Println("Задание А")
	for x := 0.11; x <= 0.36; x += 0.05 {
		fmt.Println(equation(x))
	}

	fmt.Println("Задание B")
	fmt.Println(equation(0.2))
	fmt.Println(equation(0.3))
	fmt.Println(equation(0.38))
	fmt.Println(equation(0.43))
	fmt.Println(equation(0.57))
}
