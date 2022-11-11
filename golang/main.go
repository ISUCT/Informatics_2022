package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a float64 = 0.1
		b float64 = 0.5
	)
	// Задача А, формула 17
	for x := 0.15; x <= 1.37; x += 0.25 {
		var y float64 = (a + math.Pow(math.Tan(b*x), 2)) / (b + (math.Pow(math.Pow(math.Tan(a*x), -1), 2)))
		fmt.Println(y)
	}
	// Задача Б, формула 17
	xs := [5]float64{0.2, 0.3, 0.44, 0.6, 0.56}
	for _, x := range xs {
		var y float64 = (a + math.Pow(math.Tan(b*x), 2)) / (b + (math.Pow(math.Pow(math.Tan(a*x), -1), 2)))
		fmt.Println(y)
	}
}
