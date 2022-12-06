package internal

// Задача с Четное-Нечетное
func EvenOdd(a int) string {
	if a%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// Количество овец
func Sheep(arr []bool) int {
	var result, i int
	for ; i < len(arr); i++ {
		if arr[i] && (arr[i] || !arr[i]) {
			result++
		}
	}
	return result
}

// Подсчет обезьян
func Monkey(number int) []int {
	var result []int
	for n := 1; n <= number; n++ {
		result = append(result, n)
	}
	return result
}

// Подсчет бумаги
func Printer(n, m int) int {
	if n < 0 || m < 0 {
		return 0
	}
	return n * m
}

// Драконы
func Dragons(ammo int, dragons int) bool {
	return ammo/2 >= dragons
}
