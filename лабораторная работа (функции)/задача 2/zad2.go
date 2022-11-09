package main

import (
	"fmt"
	"math"
)

func main() {

	var a = 1.6
	var znach = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}

	for _, x := range znach {
		b := math.Pow(x, 2) - 1
		y := math.Pow(a, b) - math.Log10(b) + math.Pow(b, (1/3))
		fmt.Println("x = ", x, " ", "y = ", y)
	}
}
