package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	fmt.Println(taskA(0.8, 0.4, 1.23))
	fmt.Println(taskB(0.8, 0.4))
	fmt.Println(sum_of_minimums([][]int{{1, 2, 3, 3}, {1, 2, 3, 4, 5}, {1, 3, 4, 5}}))
}

func equation(a, b, x float64) float64 {
	return ((math.Pow((x-a), 2/3) + math.Pow(math.Abs(x+b), 0.2)) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), 1/9))
}

func taskA(a, b, x float64) []float64 {
	fmt.Println("Part A")
	var array []float64
	for x <= 7.23 {
		array = append(array, equation(a, b, x))
		x = x + 1.2
	}
	return array
}

func taskB(a, b float64) []float64 {
	fmt.Println("Part B")
	var array []float64
	var x float64 = 1.88
	array = append(array, equation(a, b, x))

	x = 2.26
	array = append(array, equation(a, b, x))

	x = 3.84
	array = append(array, equation(a, b, x))

	x = 4.55
	array = append(array, equation(a, b, x))

	x = -6.21
	array = append(array, equation(a, b, x))
	return array
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	if number%2 != 0 {
		return "Odd"
	}

	return " "
}

func CountSheeps(numbers []bool) int {
	var sum int = 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] {
			sum = sum + 1
		}
	}
	return sum
}

func MonkeyCount(n int) []int {
	count := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		count = append(count, i)
	}
	return count
}

func Paperwork(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	} else {
		return n * m
	}
}

func Hero(bullets, dragons int) bool {
	if dragons <= bullets/2 {
		return true
	} else {
		return false
	}
}

func correctPolishLetters(st string) string {
	st = strings.Replace(st, "ą", "a", -1)
	st = strings.Replace(st, "ć", "c", -1)
	st = strings.Replace(st, "ę", "e", -1)
	st = strings.Replace(st, "ł", "l", -1)
	st = strings.Replace(st, "ń", "n", -1)
	st = strings.Replace(st, "ó", "o", -1)
	st = strings.Replace(st, "ś", "s", -1)
	st = strings.Replace(st, "ź", "z", -1)
	st = strings.Replace(st, "ż", "z", -1)

	return st
}

func find_all(array []int, n int) []int {
	var indexes []int
	for i := 0; i < len(array); i++ {
		if array[i] == n {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func sum_of_minimums(numbers [][]int) int {
	result := 0
	for _, i := range numbers {
		sort.Ints(i)
		result += i[0]
	}
	return result
}
