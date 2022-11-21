package main

import (
	"fmt"
	"math"
)

func main() {
	const a float64 = 2.0
	fmt.Println("___Задание А___")
	fmt.Println(TaskA(a))
	fmt.Println("___Задание В___")
	fmt.Println(TaskB(a))
}

func TaskA(a float64) []float64 {
	var x float64 = 1.2
	var taskA []float64
	for ; x < 4.2; x = x + 0.6 {
		taskA = append(taskA, formula(x, a))
	}
	return taskA
}

func TaskB(a float64) []float64 {
	var taskB []float64
	var masB = []float64{1.16, 1.32, 1.47, 1.65, 1, 93}
	for i := 0; i < len(masB); i++ {
		taskB = append(taskB, formula(masB[i], a))
	}
	return taskB
}

func formula(x float64, a float64) float64 {
	var y float64 = (math.Log10(a + x)) / ((a + x) * (a + x))
	return y
}
