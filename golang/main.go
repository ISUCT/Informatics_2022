package main

import (
	"fmt"
	"math"
)

func main() {
	// 18 Формула, задача А
	a, b := 2.5, 4.6
	for x := 1.1; x <= 3.6; x += 0.5 {
		answer := math.Pow((a+b*x), 2.5) / (1 + math.Log10(a+b*x))
		fmt.Println(answer)
	}

	// 18 Формула, задача B
	variables := [5]float64{1.2, 1.28, 1.36, 1.46, 2.35}
	for _, x := range variables {
		answer := math.Pow((a+b*x), 2.5) / (1 + math.Log10(a+b*x))
		fmt.Println(answer)
	}
}
