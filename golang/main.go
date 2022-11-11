package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 2.25
	fmt.Println("№22")
	fmt.Println("А")
	for x := 1.2; x <= 2.7; x += 0.3 {
		var first = math.Pow(a, (math.Pow(x, 2) - 1))
		var second = math.Log(math.Pow(x, 2) - 1)
		var third = math.Pow(math.Pow(x, 2)-1, 1/3)
		var result = first - second + third
		fmt.Println(result)
	}

	fmt.Println("B")
	arr := []float64{1.31, 1.39, 1.44, 1.56, 1.92}
	for _, b := range arr {
		var first = math.Pow(a, (math.Pow(b, 2) - 1))
		var second = math.Log(math.Pow(b, 2) - 1)
		var third = math.Pow(math.Pow(b, 2)-1, 1/3)
		var result = first - second + third
		fmt.Println(result)
	}
}
