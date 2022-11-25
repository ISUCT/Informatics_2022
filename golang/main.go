package main

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
