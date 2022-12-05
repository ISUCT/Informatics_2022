package main

import (
	"fmt"
	"math"
)

func main() {
	taskA()
	taskB()
}

func taskA() int {
	fmt.Println("Part A")
	var x float64 = 1.23
	var a float64 = 0.8
	var b float64 = 0.4
	for x <= 7.23 {
		fmt.Println((math.Pow((x-a), 2/3) + math.Pow(math.Abs(x+b), 0.2)) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), 1/9))
		x = x + 1.2
	}
	return 0
}

func taskB() int {
	fmt.Println("Part B")
	var a float64 = 0.8
	var b float64 = 0.4

	var x float64 = 1.88
	fmt.Println((math.Pow((x-a), 2/3) + math.Pow(math.Abs(x+b), 0.2)) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), 1/9))

	x = 2.26
	fmt.Println((math.Pow((x-a), 2/3) + math.Pow(math.Abs(x+b), 0.2)) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), 1/9))

	x = 3.84
	fmt.Println((math.Pow((x-a), 2/3) + math.Pow(math.Abs(x+b), 0.2)) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), 1/9))

	x = 4.55
	fmt.Println((math.Pow((x-a), 2/3) + math.Pow(math.Abs(x+b), 0.2)) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), 1/9))

	x = -6.21
	fmt.Println((math.Pow((x-a), 2/3) + math.Pow(math.Abs(x+b), 0.2)) / math.Pow((math.Pow(x, 2)-math.Pow((a+b), 2)), 1/9))
	return 0
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
