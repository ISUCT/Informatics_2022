package main

import (
	"fmt"
	"math"
)

func main() {
	var masB = []float64{1.16, 1.32, 1.47, 1.65, 1, 93}
	const a float64 = 2.0

	// var masA []float64 = make([]float64, 0)
	fmt.Println("___Задание А___")
	fmt.Println(TaskA(a))
	fmt.Println("___Задание В___")
	fmt.Println(TaskB(a, masB))
}

func TaskA(a float64) []float64 {
	var taskA []float64
	var nx float64 = 1.2
	var kx float64 = 4.2
	var dx float64 = 0.6

	for x := nx; x <= kx; x += dx {
		taskA = append(taskA, formula(x, a))
	}
	return taskA
}

func TaskB(b float64, masB []float64) []float64 {
	var taskB []float64
	for i := 0; i < len(masB); i++ {
		taskB = append(taskB, formula(masB[i], b))
	}
	return taskB
}

func formula(x float64, a float64) float64 {
	var y float64 = (math.Log10(a + x)) / ((a + x) * (a + x))
	return y
}
