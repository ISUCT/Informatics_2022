package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задание A")
	for x := 1.81; x <= 5.31; x += 0.7 {
		var chisl = (math.Sqrt(math.Abs(7.2 - 4.2*x)))
		var znamen = (math.Sqrt(math.Pow(math.Log(x), 3)))
		var y = (chisl / znamen)
		fmt.Println(y)
	}
	fmt.Println("Задание B")
	mass_B := []float64{2.4, 2.8, 3.9, 4.7, 3.16}
	for _, b := range mass_B {
		var chisl = (math.Sqrt(math.Abs(7.2 - 4.2*b)))
		var znamen = (math.Sqrt(math.Pow(math.Log(b), 3)))
		var y = (chisl / znamen)
		fmt.Println(y)
	}
}
