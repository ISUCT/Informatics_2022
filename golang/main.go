package main

import (
	"fmt"
	"math"

	"isuct.ru/informatics2022/internal"
)

func main() {
	//Вариант 14 (Формула)
	fmt.Println("Лабораторная работа")
	fmt.Println("Task А")
	fmt.Println(TaskA(7.2, 4.2, 1.81, 5.31, 0.7))
	fmt.Println("Task B")
	fmt.Println(TaskB(7.2, 4.2, [5]float64{2.4, 2.8, 3.9, 4.7, 3.16}))
	//CodeWars №1
	fmt.Println("CodeWars №1")
	fmt.Println(internal.EvenOrOdd(7876), ",", internal.EvenOrOdd(879))
	//CodeWars №2
	fmt.Println("CodeWars №2")
	Sheeps := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	fmt.Println(internal.SheepsCounter(Sheeps))
	//CodeWars №3
	fmt.Println("CodeWars №3")
	fmt.Println(internal.MonkeysCounter(10))
	fmt.Println(internal.MonkeysCounter(1))
	//CodeWars №4
	fmt.Println("CodeWars №4")
	fmt.Println(internal.PagesCounter(24, 10))
	fmt.Println(internal.PagesCounter(24, -1))
	//CodeWars №5
	fmt.Println("CodeWars №5")
	fmt.Println(internal.Survival(7, 3))
	fmt.Println(internal.Survival(7, 4))
	//CodeWars №6
	fmt.Println("CodeWars №6")
	fmt.Println(internal.PolishAlphabet("Jędrzej Błądziński"))
	//CodeWars №7
	fmt.Println("CodeWars №7")
	fmt.Println(internal.Find(3, []int{6, 9, 3, 4, 3, 82, 11}))
	fmt.Println(internal.Find(81, []int{6, 9, 3, 4, 3, 82, 11}))
	//CodeWars №8
	fmt.Println("CodeWars №8")
	fmt.Println(internal.SumOfMinimums([][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 34, 56, 100}}))
}

// Функция для формулы
func formula(a float64, b float64, x float64) float64 {
	return (math.Sqrt(math.Abs(a - b*x))) / (math.Sqrt(math.Pow(math.Log(x), 3)))
}

// Функция для задания А
func TaskA(a, b, xn, xk, dx float64) []float64 {
	var ArrayForTaskA []float64
	for x := xn; x <= xk; x += dx {
		ArrayForTaskA = append(ArrayForTaskA, formula(a, b, x))
	}
	return ArrayForTaskA
}

// Функция для задания В
func TaskB(a, b float64, MassTaskB [5]float64) []float64 {
	var ArrayForTaskB []float64
	for i := 0; i < len(MassTaskB); i += 1 {
		ArrayForTaskB = append(ArrayForTaskB, formula(a, b, MassTaskB[i]))
	}
	return ArrayForTaskB
}
