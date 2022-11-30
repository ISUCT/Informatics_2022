package main

import (
	"fmt"
	"math"
)

func main() {
	var masB = []float64{1.16, 1.32, 1.47, 1.65, 1, 93}
	var masA = []float64{1.2, 1.8, 2.4, 3, 3.6, 4.2}
	const a float64 = 2.0
	fmt.Println("___Задание А___")
	fmt.Println(TaskA(a, masA))
	fmt.Println("___Задание В___")
	fmt.Println(TaskB(a, masB))
}

func TaskA(a float64, masA []float64) []float64 {
	var taskA []float64
	for x := 0; x < len(masA); x++ {
		taskA = append(taskA, formula(masA[x], a))
	}
	return taskA
}

func TaskB(a float64, masB []float64) []float64 {
	var taskB []float64
	for i := 0; i < len(masB); i++ {
		taskB = append(taskB, formula(masB[i], a))
	}
	return taskB
}

func formula(x float64, a float64) float64 {
	var y float64 = (math.Log10(a + x)) / ((a + x) * (a + x))
	return y
}
