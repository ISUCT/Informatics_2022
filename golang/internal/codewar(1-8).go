package internal

import (
	"strings"
)

// "Even or Odd" №1
func EvenOrOdd(number int) string {
	if number%2 != 0 {
		return "Odd"
	} else {
		return "Even"
	}
}

// "Counting sheep..." №2
func CountSheeps(numbers []bool) int {
	var count = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == true {
			count++
		}
	}
	return count
}

// "MonkeyCount" №3
func MonkeyCount(n int) []int {
	var res []int
	for i := 1; i < n+1; i++ {
		res = append(res, i)
	}
	return res
}

// "School Paperwork" №4
func Paperwork(n, m int) int {
	if (n > 0) && (m > 0) {
		return n * m
	}
	return 0
}

// "Is he gonna survive?" #5
func Hero(bullets, dragons int) bool {
	if bullets/dragons >= 2 {
		return true
	} else {
		return false
	}
}

// "Polish alphabet" #6
func CorrectPolishLetters(str string) string {
	Replace := strings.NewReplacer("ą", "a", "ć", "c", "ę", "e", "ł", "l", "ń", "n", "ó", "o", "ś", "s", "ź", "z", "ż", "z")
	ConvText := Replace.Replace(str)
	return ConvText
}

// "Find all occurrences of an element in an array" #7
func FindAll(OurNumber int, OurArray []int) []int {
	var arr []int
	for index, Number := range OurArray {
		if Number == OurNumber {
			arr = append(arr, index)
		}
	}
	return arr
}

// "Sum of Minimums!" #8
func SumOfMinimums(OurArray [][]int) int {
	var Count int
	for _, Row := range OurArray {
		for _, Number := range Row {
			if Number < Row[0] {
				Row[0] = Number
			}
		}
		Count += Row[0]
	}
	return Count
}
