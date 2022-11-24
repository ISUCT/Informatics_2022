package main

import (
	"fmt"
	"math"
)

var a2 float64 = math.Pow(0.05, 2) // О - Оптимизация
var b2 float64 = math.Pow(0.06, 2)

func Calc(x float64) float64 {
	return math.Acos(math.Pow(x, 2)-b2) / math.Asin(math.Pow(x, 2)-a2)
}

func TaskA(from, to, step float64) {
	for i := from; i <= to; i += step {
		fmt.Println(Calc(i))
	}
}

func TaskB(values [5]float64) {
	for _, v := range values {
		fmt.Println(Calc(v))
	}
}

func main() {
	fmt.Println("Task A:")
	fmt.Println("")
	TaskA(0.2, 0.95, 0.15)

	fmt.Println("Task B:")
	fmt.Println("")
	TaskB([5]float64{0.15, 0.26, 0.37, 0.48, 0.56})
}
