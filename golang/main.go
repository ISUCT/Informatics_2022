package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	var a float64 = 2.0
	var b float64 = 0.95
	var y float64 = (1 + math.Pow(math.Log10(a/b), 2.0)) / (b - math.Pow(2.718, x/a))
	return y
}

func main() {
	// Задача А
	fmt.Println("Задача А")

	for x := 1.25; x <= 2.75; x += 0.3 {
		fmt.Println("y(", x, ") =", f(x))
	}

	// Задача Б
	fmt.Println("\nЗадача Б")
	fmt.Println(f(2.2))
	fmt.Println(f(3.78))
	fmt.Println(f(4.51))
	fmt.Println(f(6.58))
	fmt.Println(f(1.2))

}
