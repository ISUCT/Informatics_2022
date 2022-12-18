package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	const a float64 = 2.5
	const b float64 = 4.6
	fmt.Print("Задача A\n")
	for i := 1.1; i < 3.6; i += 0.5 {
		var result float64 = task(a, b, i)
		fmt.Printf("Результат вычислений = %.4f\n", result)
	}
	fmt.Print("\n")
	fmt.Print("Задача B\n")
	array := [5]float64{1.2, 1.28, 1.36, 1.46, 2.35}
	for _, element := range array {
		var result float64 = task(a, b, element)
		fmt.Printf("Результат вычислений = %.4f\n", result)
	}

	fmt.Print(EvenOrOdd(68))
	fmt.Print("\n")
	sheeps := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	fmt.Print(CountSheeps(sheeps))
	fmt.Print("\n")

	fmt.Print(monkeyCount(4))
	fmt.Print("\n")

	fmt.Print(Hero(2, 5))
	fmt.Print("\n")

	fmt.Print(correctPolishLetters("Jędrzej Błądziński"))
	fmt.Print("\n")

	arrayForOccurrences := []int{
		6, 9, 3, 4, 3, 82, 11,
	}
	fmt.Print(findAll(arrayForOccurrences, 3))
	fmt.Print("\n")

	arrayForSumMinimums := [][]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
		{20, 21, 34, 56, 100},
	}
	fmt.Print(sumOfMinimums(arrayForSumMinimums))
	fmt.Print("\n")

	fmt.Print(paperwork(5, 5))
	fmt.Print("\n")
	fmt.Print(paperwork(-5, 5))
}


func task(a, b, x float64) float64 {
	return math.Pow((a+b*x), 2.5) / (1 + math.Log(a+b*x))
}


func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}


func CountSheeps(numbers []bool) int {
	if numbers != nil {
		var answer int = 0
		for _, item := range numbers {
			if item {
				answer++
			}
		}
		return answer
	}
	return 0
}


func monkeyCount(n int) []int {
	count := make([]int, n)
	for i := 1; i <= n; i++ {
		count[i-1] = i
	}
	return count
}


func Hero(bullets, dragons int) bool {
	if bullets < dragons*2 {
		return false
	} else {
		return true
	}
}


func correctPolishLetters(str string) string {
	str = strings.Replace(str, "ą", "a", -1)
	str = strings.Replace(str, "ć", "c", -1)
	str = strings.Replace(str, "ę", "e", -1)
	str = strings.Replace(str, "ł", "l", -1)
	str = strings.Replace(str, "ń", "n", -1)
	str = strings.Replace(str, "ó", "o", -1)
	str = strings.Replace(str, "ś", "s", -1)
	str = strings.Replace(str, "ź", "z", -1)
	str = strings.Replace(str, "ż", "z", -1)
	return str
}


func findAll(array []int, n int) []int {
	var numberOccurrences int = 0
	for _, item := range array {
		if item == n {
			numberOccurrences++
		}
	}
	result := make([]int, numberOccurrences)
	var currentIndex int = 0

	for index, item := range array {
		if item == n {
			result[currentIndex] = index
			currentIndex++
		}
	}

	return result
}


func sumOfMinimums(numbers [][]int) int {
	var result int = 0
	var minimum int
	for _, i := range numbers {
		minimum = i[0]
		for _, j := range i {
			if minimum > j {
				minimum = j
			}
		}
		result += minimum
	}

	return result
}


func paperwork(n int, m int) int {
	if n < 0 || m < 0 {
		return 0
	} else {
		return n * m
	}
}
