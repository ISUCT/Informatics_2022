package main

import (
	"fmt"
)

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func CountSheeps(numbers []string) int {
	var count = 0
	for _, value := range numbers {
		if value == "true" {
			count = count + 1
		}
	}
	return count
}

func monkeyCount(n int) []int {
	var monkey []int
	monkeySlice := monkey[:]
	for i := 1; i <= n; i++ {
		monkeySlice = append(monkeySlice, i)
	}
	return monkeySlice
}

func PaperWork(n1 int, m int) int {
	if n1 > 0 && m > 0 {
		return (n1 * m)
	} else {
		return 0
	}
}

func Hero(bullets, dragons int) bool {
	if bullets >= (dragons * 2) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Лабораторная работа 1")
	var a = 1.6
	fmt.Println(zadacha1(a), zadacha2(a))
	fmt.Println("Задачи codewars")
	var number = 4
	fmt.Println(EvenOrOdd(number))
	numbers := [5]string{"true", "true", "false", "true", "true"}
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
