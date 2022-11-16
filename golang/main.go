package main

import (
	"fmt"
	"math"
)

func equation(x float64) float64 {
	var a float64 = 0.8
	var b float64 = 0.4
	return (math.Pow((x-a), 2/3) + math.Pow(math.Abs(x+b), 0.2)) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), 1/9)
}

func main() {
	fmt.Println("Part A")
	var x float64 = 1.23
	for x <= 7.23 {
		fmt.Println(equation(x))
		x = x + 1.2
	}

	fmt.Println("Part B")

	x = 1.88
	fmt.Println(equation(x))

	x = 2.26
	fmt.Println(equation(x))

	x = 3.84
	fmt.Println(equation(x))

	x = 4.55
	fmt.Println(equation(x))

	x = -6.21
	fmt.Println(equation(x))
}
