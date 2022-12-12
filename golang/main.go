package main

import (
	"fmt"
	"math"
)

func main() {
	// formula 11  y = arcsin(x^a) + arccos(x^b)
	Formula11Range(2.0, 3.0, 0.11, 0.36, 0.05)
	fmt.Println("-----------------------")
	Formula11Array(2.0, 3.0, []float64{0.08, 0.26, 0.35, 0.41, 0.53})
}

func Formula11Range(a, b, xn, xk, dx float64) {
	for x := xn; x <= xk; x += dx {
		fmt.Println(math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b)))
	}
}

func Formula11Array(a, b float64, arr []float64) {
	for _, x := range arr {
		fmt.Println(math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b)))
	}
}
