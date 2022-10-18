package main

import (
	"fmt"
	"math"
)

func Log(base float64, num float64) float64 {
	return math.Log(num) / math.Log(base)
}

func problem(x float64, a float64, b float64) float64 {
	var numerator float64 = a*math.Pow(x, 1/3.0) - b*Log(5, x)
	var denominator float64 = math.Pow(math.Log10(x-1), 3)
	return numerator / denominator
}

func taskA(x_start float64, x_end float64, delta float64) []float64 {
	var valuesCount int = int(math.Floor((x_end - x_start) / delta))
	var values = make([]float64, valuesCount)

	for i := 0; i < valuesCount; i++ {
		values[i] = problem(x_start+float64(i)*delta, 4.1, 2.7)
	}

	return values
}

func taskB(args ...float64) []float64 {
	var values = make([]float64, len(args))

	for i := 0; i < len(args); i++ {
		values[i] = problem(args[i], 4.1, 2.7)
	}
	return values
}

func main() {
	fmt.Println("задание А", taskA(1.5, 3.5, 0.4))
	fmt.Println("Задание Б", taskB(1.9, 2.15, 2.34, 2.74, 3.16))
}
