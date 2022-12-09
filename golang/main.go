package main

import (
	"fmt"
	"math"

	"isuct.ru/informatics2022/internal"
)

func main() {
	//Лабораторная работа
	fmt.Println("Task А")
	fmt.Println(SolveTaskA(4.1, 2.7, 1.2, 5.2, 0.8))

	fmt.Println("Task B")
	MassTaskB := [5]float64{1.9, 2.15, 2.34, 2.73, 3.16}
	fmt.Println(SolveTaskB(4.1, 2.7, MassTaskB))
	//CodeWars 1
	fmt.Println("CodeWarsTask 1")
	fmt.Println(internal.CheckOddOrEven(777), ",", internal.CheckOddOrEven(666))
	//CodeWars 2
	fmt.Println("CodeWarsTask 2")
	Sheep := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	fmt.Println(internal.SheepsCounter(Sheep))
	//CodeWars 3
	fmt.Println("CodeWarsTask 3")
	fmt.Println(internal.MonkeysCounter(10))
	fmt.Println(internal.MonkeysCounter(1))
	//CodeWars 4
	fmt.Println("CodeWarsTask 4")
	fmt.Println(internal.CountPaperwork(5, 5))
	fmt.Println(internal.CountPaperwork(-5, 5))
	//CodeWars 5
	fmt.Println("CodeWarsTask 5")
	fmt.Println(internal.IsHeGonnaSurvive(5, 2))
	fmt.Println(internal.IsHeGonnaSurvive(1, 2))
	//CodeWars 6
	fmt.Println("CodeWarsTask 6")
	fmt.Println(internal.ReplacePolishLetters("Jędrzej Błądziński"))
	//CodeWars 7
	fmt.Println("CodeWarsTask 7")
	FindAllExampleArray := []int{6, 9, 3, 4, 3, 82, 11}
	fmt.Println(internal.FindAll(3, FindAllExampleArray))
	fmt.Println(internal.FindAll(1, FindAllExampleArray))
	//CodeWars 8
	fmt.Println("CodeWarsTask 8")
	SumOfMinimumsExampleArray := [][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 34, 56, 100}}
	fmt.Println(internal.SumOfMinimums(SumOfMinimumsExampleArray))
}

// Function For The Formula
func formula(a float64, b float64, x float64) float64 {
	return ((a * math.Sqrt(x)) - (b * (math.Log2(x) / math.Log2(5)))) / (math.Log(math.Abs(x - 1)))
}

// Funtion For The Task A
func SolveTaskA(a float64, b float64, xn float64, xk float64, dx float64) []float64 {
	var ArrayForTaskA []float64
	for x := xn; x <= xk; x += dx {
		ArrayForTaskA = append(ArrayForTaskA, formula(a, b, x))
	}
	return ArrayForTaskA
}

// Funtion For The Task B
func SolveTaskB(a float64, b float64, MassTaskB [5]float64) []float64 {
	var ArrayForTaskB []float64
	for i := 0; i < len(MassTaskB); i += 1 {
		ArrayForTaskB = append(ArrayForTaskB, formula(a, b, MassTaskB[i]))
	}
	return ArrayForTaskB
}
