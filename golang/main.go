package main

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

// Поменяла monkeyCount на MonkeyCount, так как иначе линтер ругался на неиспользованную функцию
func MonkeyCount(n int) []int {
	counting := make([]int, 0, n)
	for i := 1; i <= n; i += 1 {
		counting = append(counting, i)
	}
	return counting
}

// Это задание нельзя было выполнить на Go на Codewars, поэтому написала для него тесты в main_test.go
func PaperWork(n, m int) int {
	if n < 0 || m < 0 {
		return 0
	} else {
		return n * m
	}
}

func Hero(bullets, dragons int) bool {
	return dragons*2 <= bullets
}

func main() {
}
