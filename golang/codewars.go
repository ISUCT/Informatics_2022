package main

import (
	"strings"
)

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func CountSheeps(numbers []bool) int {
	var count = 0
	for _, value := range numbers {
		if value {
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

func correctPolishLetters(TextPolish string) string {
	replacer := strings.NewReplacer("ą", "a", "ć", "c", "ę", "e", "ł", "l", "ń", "n", "ó", "o", "ś", "s", "ź", "z", "ż", "z")
	TextRus := replacer.Replace(TextPolish)
	return TextRus
}

func findAll(n int, massiv []int) []int {
	var out []int
	for i, num := range massiv {
		if num == n {
			out = append(out, i)
		}
	}
	return out
}

func min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

func sumOfMinimums(matrix [][]int) int {
	var summa = 0
	for i := 0; i <= len(matrix); i++ {
		for j := 0; j <= len(matrix); j++ {
			if matrix[i][j] >= matrix[i][j+1] {
				summa += matrix[i][j+1]
			}
		}
	}
	return summa
}
