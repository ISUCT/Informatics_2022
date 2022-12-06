package main

import (
	"fmt"
	"math"

	"isuct.ru/informatics2022/internal"
)

func main() {
	var masB = []float64{1.16, 1.32, 1.47, 1.65, 1, 93}
	const a float64 = 2.0
	var nx float64 = 1.2
	var kx float64 = 4.2
	var dx float64 = 0.6

	fmt.Println("___Задание А___")
	fmt.Println(TaskA(a, nx, kx, dx))
	fmt.Println("___Задание В___")
	fmt.Println(TaskB(a, masB))

	// Задача с Четное-Нечетное
	fmt.Println(internal.EvenOdd(1703))
	// Количество овец
	fmt.Println(internal.Sheep([]bool{true, false, true, true, true, true, false, false}))
	// Подсчет обезьян
	fmt.Println(internal.Monkey(10))
	// Подсчет бумаги
	fmt.Println(internal.Printer(5, 2))
	// Драконы
	fmt.Println(internal.Dragons(5, 2))
}

func TaskA(a, nx, kx, dx float64) []float64 {
	var taskA []float64

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
