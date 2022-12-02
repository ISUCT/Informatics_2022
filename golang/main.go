package main

import (
	"fmt"
)

func main() {
	fmt.Println("Лабораторная работа 1")
	var a = 1.6
	fmt.Println("Задание А ")
	var xn = 1.2
	var xk = 4.2
	var dx = 0.6
	fmt.Println("y = ", zadacha1(a, xn, xk, dx))
	fmt.Println("Задание В ")
	var znach = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	fmt.Println("y = ", zadacha2(a, znach[:]))
	fmt.Println("Задачи codewars")
	fmt.Println("EvenOrOdd")
	var number = 4
	fmt.Println(EvenOrOdd(number))
	fmt.Println("CountSheeps")
	numbers := [5]bool{true, true, false, true, true}
	fmt.Println(CountSheeps(numbers[:]))
	fmt.Println("MonkeyCount")
	var n = 5
	fmt.Println(monkeyCount(n))
	fmt.Println("PaperWork")
	var n1 = 5
	var m = 4
	fmt.Println(PaperWork(n1, m))
	fmt.Println("Hero")
	var bullets = 5
	var dragons = 4
	fmt.Println(Hero(bullets, dragons))
	fmt.Println("CorrectPolishLetters")
	TextPolish := "Jędrzej Błądziński"
	fmt.Println(correctPolishLetters(TextPolish))
	fmt.Println("FindAll")
	var n2 = 3
	massiv := [5]int{3, 4, 3, 5, 7}
	fmt.Println(findAll(n2, massiv[:]))
	fmt.Println("SumOfMinimums")
	var matrix = [][]uint{{1, 4, 8}, {5, 7, 4}, {4, 9, 2}}
	fmt.Println(sumOfMinimums(matrix[:][:]))
}
