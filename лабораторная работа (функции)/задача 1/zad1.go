package main

import (
	"fmt"
	"math"
)

func main() {

	var a = 1.6

	for x := 1.2; x <= 3.7; x += 0.5 {
		b := math.Pow(x, 2) - 1
		y := math.Pow(a, b) - math.Log10(b) + math.Pow(b, (1/3))
		fmt.Println("x = ", x, " ", "y = ", y)
	}
}
