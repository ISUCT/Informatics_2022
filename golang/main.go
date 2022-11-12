package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Задание А")
	for x := 1.23; x <= 7.23; x += 1.2 {
		var ch = (math.Pow(math.Pow(x-0.8, 2.0), 1/3.0) + (math.Pow(math.Abs(x+0.4), 1/5.0)))
		var zn = (math.Pow(math.Pow(x, 2.0)-math.Pow(0.8+0.4, 2.0), 1/9.0))
		var y = (ch / zn)
		fmt.Println(y)
	}

	fmt.Println("Задание B")
	mass_B := []float64{1.9, 2.15, 2.34, 2.73, 3.16}
	for _, x := range mass_B {
		var ch = (math.Pow(math.Pow(x-0.8, 2.0), 1/3.0) + (math.Pow(math.Abs(x+0.4), 1/5.0)))
		var zn = (math.Pow(math.Pow(x, 2.0)-math.Pow(0.8+0.4, 2.0), 1/9.0))
		var y = (ch / zn)
		fmt.Println(y)

	}
}
