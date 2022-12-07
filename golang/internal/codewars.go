package internal

import "sort"

func EvenOrOdd(x int) (answer string) {
	if (x % 2) == 0 {
		answer = "Even"
	} else {
		answer = "Odd"
	}
	return
}

func Sheep(sheep []bool) (answer int) {
	answer = 0
	for _, value := range sheep {
		if value {
			answer += 1
		}
	}
	return
}

func CountWithSon(number int) (answer []int) {
	for i := 1; i <= number; i++ {
		answer = append(answer, i)
	}
	return
}

func School(n, m int) (answer int) {
	if (n < 0) || (m < 0) {
		answer = 0
	} else {
		answer = n * m
	}
	return
}

func LuckyHero(bullets, dragons int) (answer bool) {
	if float32(bullets/2.0) >= float32(dragons) {
		answer = true
	} else {
		answer = false
	}
	return
}

func PolToEng(message string) (answer string) {
	var alphabet = map[string]string{
		"ą": "a",
		"ć": "c",
		"ę": "e",
		"ł": "l",
		"ń": "n",
		"ó": "o",
		"ś": "s",
		"ź": "z",
		"ż": "z",
	}
	for _, x := range message {
		if alphabet[string(x)] != "" {
			answer += alphabet[string(x)]
		} else {
			answer += string(x)
		}
	}
	return
}

func FindAll(arr []int, k int) (answer []int) {
	for index, value := range arr {
		if value == k {
			answer = append(answer, index)
		}
	}
	return
}

func OnlyMin(arr [][]int) (answer int) {
	for _, n_1 := range arr {
		sort.Ints(n_1[:])
		answer += n_1[0]
	}
	return answer
}
