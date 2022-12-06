package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(taskA(0.8, 0.4, 1.23))
	fmt.Println(taskB(0.8, 0.4))
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
