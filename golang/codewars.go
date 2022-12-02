package main

import "strings"

// ===== Часть 1 =====

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func CountSheeps(numbers []bool) int {
	count := 0
	for _, elem := range numbers {
		if elem {
			count += 1
		}
	}
	return count
}

func MonkeyCount(n int) []int {
	counting := make([]int, 0, n)
	for i := 1; i <= n; i += 1 {
		counting = append(counting, i)
	}
	return counting
}

func Paperwork(n, m int) int {
	if n < 0 || m < 0 {
		return 0
	} else {
		return n * m
	}
}

func Hero(bullets, dragons int) bool {
	return dragons*2 <= bullets
}

// ===== Часть 2 =====

func SumOfMinimums(numbers [][]uint) uint {
	var sum uint
	for _, array := range numbers {
		for _, element := range array {
			if element < array[0] {
				array[0] = element
			}
		}
		sum += array[0]
	}
	return sum
}

func FindAll(numbers []int, someNumber int) []int {
	indexes := make([]int, 0, len(numbers))
	for index, element := range numbers {
		if element == someNumber {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

func CorrectPolishLetters(str string) string {
	multipleReplaces := strings.
		NewReplacer(
			"ą", "a",
			"ć", "c",
			"ę", "e",
			"ł", "l",
			"ń", "n",
			"ó", "o",
			"ś", "s",
			"ź", "z",
			"ż", "z",
		)
	modifiedString := multipleReplaces.Replace(str)
	return modifiedString
}
