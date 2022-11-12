package main

import (
	"fmt"
	"math"
)

// variant 12

func CalculateEquation(x float64, a float64) float64 {
	var temp float64 = math.Pow(x, 2) - 1
	return math.Pow(a, temp) - math.Log10(temp) + math.Pow(temp, float64(1)/3)
}

func TaskA(startX float64, endX float64, deltaX float64, a float64) []float64 {
	lenth := int32((endX-startX)/deltaX) + 1
	arrResults := make([]float64, lenth)
	numRes := 0
	for i := startX; i <= endX; i += deltaX {
		arrResults[numRes] = CalculateEquation(i, a)
		numRes++
	}
	return arrResults
}

func PrintResTaskA(startX float64, endX float64, deltaX float64, a float64) {
	res := TaskA(startX, endX, deltaX, a)
	for i := 0; i < len(res); i++ {
		fmt.Println("X = ", startX+deltaX*float64(i), "\tY = ", res[i])
	}
}

func TaskB(x []float64, a float64) []float64 {
	lenth := len(x)
	arrResults := make([]float64, lenth)
	for i := 0; i < lenth; i++ {
		arrResults[i] = CalculateEquation(x[i], a)
	}
	return arrResults
}

func PrintResTaskB(x []float64, a float64) {
	res := TaskB(x, a)
	for i := 0; i < len(res); i++ {
		fmt.Println("x = ", x[i], "\ty = ", res[i])
	}
}

func main() {
	var a float64 = 1.6
	startX := 1.2
	endX := 3.7
	deltaX := 0.5
	x := []float64{1.28, 1.36, 2.47, 3.68, 4.56}

	// Task A
	fmt.Println("Task A")
	PrintResTaskA(startX, endX, deltaX, a)

	// Task B
	fmt.Println("Task B")
	PrintResTaskB(x, a)
}
