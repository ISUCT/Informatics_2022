package main

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func CountSheeps(numbers []string) int {
	var count = 0
	for _, value := range numbers {
		if value == "true" {
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
