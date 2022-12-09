package internal

import "strings"

// Функция для CodeWars №1
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// Функция для CodeWars №2
func SheepsCounter(SheepsTrue []bool) int {
	var n = 0
	for i := 0; i < len(SheepsTrue); i++ {
		if SheepsTrue[i] {
			n += 1
		}
	}
	return n
}

// Функция для CodeWars №3
func MonkeysCounter(Amount int) []int {
	k := []int{}
	for i := 1; i <= Amount; i++ {
		k = append(k, i)
	}
	return k
}

// Функция для CodeWars №4
func PagesCounter(Classmate, Page int) int {
	if Classmate < 0 || Page < 0 {
		return 0
	}
	return Classmate * Page
}

// Функция для CodeWars №5
func Survival(Bullets, Dragons int) bool {
	return Bullets >= Dragons*2
}

// Функция для CodeWars №6
func PolishAlphabet(PText string) string {
	ReplaceAlphabet := strings.NewReplacer("ą", "a", "ć", "c", "ę", "e", "ł", "l", "ń", "n", "ó", "o", "ś", "s", "ź", "z", "ż", "z")
	CovertedText := ReplaceAlphabet.Replace(PText)
	return CovertedText
}

// Функция для CodeWars №7
func Find(FirstNumber int, FirstArray []int) []int {
	var SecondArray []int
	for index, SecondNumber := range FirstArray {
		if SecondNumber == FirstNumber {
			SecondArray = append(SecondArray, index)
		}
	}
	return SecondArray
}

// Функция для CodeWars №8
func SumOfMinimums(OurArray [][]int) int {
	var Amount int
	for _, Row := range OurArray {
		for _, Number := range Row {
			if Number < Row[0] {
				Row[0] = Number
			}
		}
		Amount += Row[0]
	}
	return Amount
}
