package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 1.35
	var b = 0.98
	fmt.Println("Задане a")
	for x := 1.14; x <= 4.24; x += 0.62 {
		y := (math.Pow((a*x + b), 1/3)) / (math.Pow(math.Log10(x), 2))
		fmt.Println("x = ", x, " ", "y = ", y)
	}

	var znach = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
	fmt.Println("Задане b")
	for _, x1 := range znach {
		y1 := (math.Pow((a*x1 + b), 1/3)) / (math.Pow(math.Log10(x1), 2))
		fmt.Println("x = ", x1, " ", "y = ", y1)

	}
}
