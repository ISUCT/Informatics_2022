package internal

import (
	"fmt"
	"math"
	"strings"
)

// Задание 1
func EvenOrOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// Задание 2
func CountingSheep(Ovces []bool) {
	var c int64 = 0
	for i := 0; i < len(Ovces); i++ {
		if Ovces[i] == true {
			c++
		}
	}
	fmt.Println("Количество польских овец ", c)
}

// Задание 3
func MonkeysCount(n int) {
	var monkey = []int{}
	for i := 1; i <= n; i++ {
		monkey = append(monkey, int(i))
	}
	fmt.Println(monkey)
}

// Задание 4
func PaperworkCount(n int, m int) int {
	if m < 0 || n < 0 {
		return 0
	} else {
		return m * n
	}
}

// Задание 5
func HeroWithGunShootsDragons(d int, b int) string {
	if (d*2)-b <= 0 {
		return "Hero Win"
	} else {
		return "Hero Died"
	}
}

// Задание 6
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

func FindAll(arr []int, search int) []int {
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
