package main

import (
	"fmt"
	"math"
)

func Log(base float64, num float64) float64 {
	return math.Log(num) / math.Log(base)
}

func problem(x float64) float64 {
	var a float64 = 4.1
	var b float64 = 2.7
	var numerator float64 = a*math.Pow(x, 1/3) - b*Log(5, x)
	var denominator float64 = math.Pow(math.Log10(x-1), 3)
	return numerator / denominator
}

func taskA() []float64 {
	var x_start float64 = 1.5
	var x_end float64 = 3.5
	var delta float64 = 0.4
	var valuesCount int = int(math.Floor((x_end - x_start) / delta))
	var values = make([]float64, valuesCount)

	for i := 0; i < valuesCount; i++ {
		values[i] = problem(x_start + float64(i)*delta)
	}

	return values
}

func taskB() [5]float64 {
	var nums = [...]float64{1.9, 2.15, 2.34, 2.74, 3.16}
	var values [5]float64

	for i := 0; i < len(nums); i++ {
		values[i] = problem(nums[i])
	}
	return values
}

func main() {
	fmt.Println("задание А", taskA())
	fmt.Println("Задание Б", taskB())
}
