package main

import (
	"fmt"
	"math"
)

func function(a, x float64) float64 {
	var y float64 = math.Pow(a, (math.Pow(x, 2)-1)) - math.Log10(math.Pow(x, 2)-1) + math.Pow((math.Pow(x, 2)-1), (1.0/3.0))
	return y
}

func main() {
	var a = 1.6
	fmt.Println("Задание А")
	for x := 1.2; x <= 3.7; x += 0.5 {
		fmt.Println("x = ", x, " ", "y = ", function(a, x))
	}
	var znach = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	fmt.Println("Задание B")
	for _, x1 := range znach {
		fmt.Println("x = ", x1, " ", "y = ", function(a, x1))
	}
}
