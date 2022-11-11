package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задание А")
	for x := 1.2; x <= 5.2; x += 0.8 {
		var chisl = ((4.1 * math.Sqrt(x)) - (2.7 * (math.Log2(x) / math.Log2(5))))
		var znamen = (math.Log(math.Abs(x - 1)))
		var y = (chisl / znamen)
		fmt.Println(y)
	}

	fmt.Println("Задание B")
	mass_B := []float64{1.9, 2.15, 2.34, 2.73, 3.16}
	for _, b := range mass_B {
		var chisl = ((4.1 * math.Sqrt(b)) - (2.7 * (math.Log2(b) / math.Log2(5))))
		var znamen = (math.Log(math.Abs(b - 1)))
		var y = (chisl / znamen)
		fmt.Println(y)

	}
}
