package main

import (
	"fmt"
	"math"
)

//Lab 24 Chukhlova V 1.42

func main() {
	var a = 7.2
	var b = 1.3
	fmt.Println("Задание А")
	for x := 1.56; x <= 4.71; x += 0.63 {
		c := (a + (b * x)) / (math.Pow(math.Log10(x), 3))
		y := math.Pow(c, (1 / 5.0))
		fmt.Println("x = ", x, " ", "y = ", y)
	}
	var mass = [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}
	fmt.Println("Задание B")
	for _, x1 := range mass {
		c := (a + (b * x1)) / (math.Pow(math.Log10(x1), 3))
		y1 := math.Pow(c, (1 / 5.0))
		fmt.Println("x = ", x1, " ", "y = ", y1)
	}
}
