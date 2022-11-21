package main

import (
	"fmt"
	"math"
)

func main() {
	const a float64 = 2.0
	const b float64 = 0.95
	fmt.Println("Задание А")
	fmt.Println(TaskA(a, b))
	fmt.Println("Задание Б")
	fmt.Println(TaskB(a, b))
}

func TaskA(a, b float64) []float64 {
	var x float64 = 1.25
	var taskA []float64
	for ; x <= 2.75; x += 0.3 {
		taskA = append(taskA, formula(x, a, b))
	}
	return taskA
}

func TaskB(a, b float64) []float64 {
	var taskB []float64
	var array = []float64{2.2, 3.78, 4.51, 6.58, 1.2}
	for i := 0; i < len(array); i++ {
		taskB = append(taskB, formula(array[i], a, b))
	}
	return taskB
}

func formula(x, a, b float64) float64 {
	var y float64 = (1 + math.Pow(math.Log10(a/b), 2.0)) / (b - math.Pow(2.718, x/a))
	return y
}
