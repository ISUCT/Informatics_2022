package main

import (
	"fmt"
	"math"
)

func main() {
	b := 2.5
	// Задача а формула 1
	for x := 1.28; x <= 3.28; x += 0.4 {
		var y float64 = (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Pow(math.Pow(b, 3)+math.Pow(x, 3), 1/3))
		fmt.Println(y)
	}
	// Задача б формула 1
	L := [5]float64{1.1, 2.4, 3.6, 1.7, 3.9}
	for _, x := range L {
		var y float64 = (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / (math.Pow(math.Pow(b, 3)+math.Pow(x, 3), 1/3))
		fmt.Println(y)
	}
}
