package main

import "strings"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func CountingSheep(count []bool) int {
	sheeps := 0
	for i := 0; i < len(count); i++ {
		if count[i] {
			sheeps ++
		}
	}
	return sheeps
}

func CountTheMonkeys(n int) []int {
	count := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		count = append(count, i)
	}
	return count
}

func SchoolPaperwork(n, m int) int {
	if n < 0 || m < 0 {
		return 0
	} else {
		return n * m
	}
}

func IsHeGonnaSurvive(bullets, dragons int) string {
	if bullets/2 >= dragons {
		return "True"
	} else {
		return "False"
	}
}


func Polish_alphabet(PolishText string) string {
	translator := strings.NewReplacer(
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
	return translator.Replace(PolishText)
}

func Find_all_occ(array []int, n int) []int {
	index := make([]int, 0, len(array))
	for i, value := range array {
		if value == n {
			index = append(index, i)
		}
	}
	return index
}

func Sum_of_Minimums(array [][]uint) uint {
	var sum uint
	for _, nest_array := range array {
		for _, minimum := range nest_array {
			if minimum < nest_array[0] {
				nest_array[0] = minimum
			} 
		}
		sum += nest_array[0]
	}
	return sum
}
