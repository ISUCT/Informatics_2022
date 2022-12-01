package internal

import "math"

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

// I hate linter, because it can't import "strings" and add string.Replace function
// func PolishCow(text string) string {
// 	var symbols = map[string]string{
// 		"ą": "a",
// 		"ć": "c",
// 		"ę": "e",
// 		"ł": "l",
// 		"ń": "n",
// 		"ó": "o",
// 		"ś": "s",
// 		"ź": "z",
// 		"ż": "z",
// 	}
// 	var symbolsValues = []string{
// 		"a", "c", "e", "l", "n", "o", "s", "z",
// 	}
// 	for i := 0; i < len(symbolsValues); i++ {
// 		text = text.Replace(symbolsValues[i], )
// 	}
// 	return text
// }

func Find_all(arr []string, search string) []int {
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
