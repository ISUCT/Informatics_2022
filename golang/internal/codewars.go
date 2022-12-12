package internal

import (
	"math"
	"strings"
)

// Поольский Алфавит
func PolisAlphabet(text string) string {
	var symbols = [][]string{
		{"ą", "a"},
		{"ć", "c"},
		{"ę", "e"},
		{"ł", "l"},
		{"ń", "n"},
		{"ó", "o"},
		{"ź", "z"},
		{"ż", "z"},
	}
	for i := 0; i < len(symbols); i++ {
		text = strings.Replace(text, symbols[i][0], symbols[i][1], -1)
	}
	return text
}

// Find All
func FindAll(arr []int, number int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		if arr[i] == number {
			res = append(res, i)
		}
	}
	return res
}

// Sum Min
func SumMin(arr [][]int, m, n int) int {
	var res int
	for i := 0; i < m; i++ {
		min := math.Inf(1)
		for a := 0; a < n; a++ {
			if float64(arr[i][a]) < min {
				min = float64(arr[i][a])
			}
		}
		res = res + int(min)
	}
	return res
}
