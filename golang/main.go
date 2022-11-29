package main

import (
	"fmt"
)

func main() {
	fmt.Println("Лабораторная работа 1")
	var a = 1.6
	fmt.Println("Задание А ")
	fmt.Println("y = ", zadacha1(a))
	fmt.Println("Задание В ")
	fmt.Println("y = ", zadacha2(a))
	fmt.Println("Задачи codewars")
	fmt.Println("Задача 1")
	var number = 4
	fmt.Println(EvenOrOdd(number))
	fmt.Println("Задача 2")
	numbers := [5]bool{true, true, false, true, true}
	fmt.Println(CountSheeps(numbers[:]))
	fmt.Println("Задача 3")
	var n = 5
	fmt.Println(monkeyCount(n))
	fmt.Println("Задача 4")
	var n1 = 5
	var m = 4
	fmt.Println(PaperWork(n1, m))
	fmt.Println("Задача 5")
	var bullets = 5
	var dragons = 4
	fmt.Println(Hero(bullets, dragons))
}
