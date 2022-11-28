package main

import (
	"fmt"
)

func main() {
	fmt.Println("Лабораторная работа 1")
	var a = 1.6
	fmt.Println("Задание А ")
	for x := 1.2; x <= 3.7; x += 0.5 {
		fmt.Println(zadacha1(a, x))
	}
	fmt.Println("Задание В ")
	var znach = [5]float64{1.28, 1.36, 2.47, 3.68, 4.56}
	for _, x1 := range znach {
		fmt.Println(zadacha2(a, x1))
	}
	fmt.Println("Задачи codewars")
	var number = 4
	fmt.Println(EvenOrOdd(number))
	numbers := [5]bool{true, true, false, true, true}
	fmt.Println(CountSheeps(numbers[:]))
	var n = 5
	fmt.Println(monkeyCount(n))
	var n1 = 5
	var m = 4
	fmt.Println(PaperWork(n1, m))
	var bullets = 5
	var dragons = 4
	fmt.Println(Hero(bullets, dragons))
}
