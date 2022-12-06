package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var a float64 = 1.1
	var b float64 = 0.09

	fmt.Printf("Задача А\n")
	for i := 1.2; i < 2.2; i += 0.2 {
		Problem(a, b, i)
	}
	fmt.Printf("-------------------------------\n")

	fmt.Printf("Задача B\n")

	xValues := [5]float64{1.21, 1.76, 2.53, 3.48, 4.52}
	for _, element := range xValues {
		Problem(a, b, element)
	}

	fmt.Printf("-------------------------------\n")
	fmt.Printf("Codewars section\n")
	fmt.Printf("-------------------------------\n")

	fmt.Printf("Is hero survive?\n") // 1
	fmt.Printf("bullets = 10, dragons = 5, result = %t\n", Hero(10, 5))
	fmt.Printf("\n")

	fmt.Printf("School Paperwork\n") // 2
	fmt.Printf("n = 5, m = 5, result = %d\n", paperwork(5, 5))
	fmt.Printf("n = -5, m = 5, result = %d\n", paperwork(-5, 5))
	fmt.Printf("\n")

	fmt.Printf("Count the Monkeys\n") // 3
	fmt.Printf("n = 10, result = %v\n", monkeyCount(10))
	fmt.Printf("\n")

	fmt.Printf("Counting sheep\n") // 4
	sheeps := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	fmt.Printf("sheeps = %v, result = %v\n", sheeps, CountSheeps(sheeps))
	fmt.Printf("\n")

	fmt.Printf("Even or Odd\n") // 5
	fmt.Printf("number = 2, result = %s\n", EvenOrOdd(2))
	fmt.Printf("number = 3, result = %s\n", EvenOrOdd(3))
	fmt.Printf("\n")

	fmt.Printf("Sum of Minimums\n") // 6
	arrayForSumMinimums := [][]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
		{20, 21, 34, 56, 100},
	}
	fmt.Printf("numbers = %v, result = %d\n", arrayForSumMinimums, sumOfMinimums(arrayForSumMinimums))
	fmt.Printf("\n")

	fmt.Printf("Find all occurrences of an element in an array\n") // 7
	arrayForOccurrences := []int{
		6, 9, 3, 4, 3, 82, 11,
	}
	fmt.Printf("numbers = %v, n = 3, result = %d\n", arrayForOccurrences, findAll(arrayForOccurrences, 3))
	fmt.Printf("\n")

	fmt.Printf("Polish alphabet\n") // 7
	fmt.Printf("input = Jędrzej Błądziński, result = %s\n", correctPolishLetters("Jędrzej Błądziński"))
}

func Problem(a, b, x float64) {
	var result float64 = math.Log(x*x-1) / Log(5, (a*x*x-b))
	fmt.Printf("a = %.4f, b = %.4f, x = %.4f, result = %.4f\n", a, b, x, result)
}

func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

// CountSheeps
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

// EvenOrOdd
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// monkeyCount
func monkeyCount(n int) []int {
	count := make([]int, n)
	for i := 1; i <= n; i++ {
		count[i-1] = i
	}
	return count
}

// Hero and dragons
func Hero(bullets, dragons int) bool {
	if bullets < dragons*2 {
		return false
	} else {
		return true
	}
}

// correctPolishLetters
func correctPolishLetters(str string) string {
	str = strings.Replace(str, "ą", "e", -1)
	str = strings.Replace(str, "ć", "e", -1)
	str = strings.Replace(str, "ę", "e", -1)
	str = strings.Replace(str, "ł", "e", -1)
	str = strings.Replace(str, "ń", "e", -1)
	str = strings.Replace(str, "ó", "e", -1)
	str = strings.Replace(str, "ś", "e", -1)
	str = strings.Replace(str, "ź", "e", -1)
	str = strings.Replace(str, "ż", "e", -1)
	return str
}

// numberOccurrences
func numberOccurrences(array []int, n int) int {
	var result int = 0
	for _, item := range array {
		if item == n {
			result++
		}
	}
	return result
}

func findAll(array []int, n int) []int {
	result := make([]int, numberOccurrences(array, n))
	var currentIndex int = 0

	for index, item := range array {
		if item == n {
			result[currentIndex] = index
			currentIndex++
		}
	}

	return result
}

// sumOfMinimums
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

// School Paperwork
func paperwork(n int, m int) int {
	if n < 0 || m < 0 {
		return 0
	} else {
		return n * m
	}
}
