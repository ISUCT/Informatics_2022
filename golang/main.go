package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 1.6
	fmt.Println("Задание А")
	for x := 1.2; x <= 3.7; x += 0.5 {
		b := math.Pow(x, 2) - 1
		y := math.Pow(a, b) - math.Log10(b) + math.Pow(b, (1/3))

		fmt.Println("x = ", x, " ", "y = ", y)
	}
	var znach = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	fmt.Println("Задание B")
	for _, x1 := range znach {
		b := math.Pow(x1, 2) - 1
		y1 := math.Pow(a, b) - math.Log10(b) + math.Pow(b, (1/3))
		fmt.Println("x = ", x1, " ", "y = ", y1)
	}
}
