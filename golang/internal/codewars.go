package internal

import (
	"math"
	"strings"
)

func EvenOrOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func CountingSheep(arr []bool) int {
	var res int
	var i int
	for ; i < len(arr); i++ {
		if arr[i] && (arr[i] || !arr[i]) {
			res++
		}
	}
	return res
}

func MonkeysCount(num int) []int {
	var res []int
	for n := 1; n <= num; n++ {
		res = append(res, n)
	}
	return res
}

func PaperworkCount(n int, m int) int {
	if n < 0 || m < 0 {
		return 0
	}
	return n * m
}

func HeroWithGunShootsDragons(ammo int, dragons int) bool {
	return ammo/2 >= dragons
}

//Codewars 2: Возвращение

func PolishCow(text string) string {
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

func FindAll(arr []string, search string) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			res = append(res, i)
		}
	}
	return res
}

func SumOfMin(arr [][]int, m int, n int) int {
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
