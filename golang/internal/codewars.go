package internal

import "strings"

// SolveCodeWars 1
func CheckOddOrEven(num int) string {
	if num%2 != 0 {
		return "Odd"
	} else {
		return "Even"
	}
}

// SolveCodeWars 2
func SheepsCounter(SheepsArray []bool) int {
	var NumberOfSheeps = 0
	for i := 0; i < len(SheepsArray); i += 1 {
		if SheepsArray[i] {
			NumberOfSheeps += 1
		}
	}
	return NumberOfSheeps
}

// SolveCodeWars 3
func MonkeysCounter(Amount int) []int {
	AmountArray := []int{}
	for i := 1; i <= Amount; i += 1 {
		AmountArray = append(AmountArray, i)
	}
	return AmountArray
}

// SolveCodeWars 4
func CountPaperwork(Classmates int, Paperwork int) int {
	if Classmates < 0 || Paperwork < 0 {
		return 0
	}
	return Classmates * Paperwork
}

// SolveCodeWars 5
func IsHeGonnaSurvive(BulletsCount int, DragonsCount int) bool {
	return BulletsCount >= DragonsCount*2
}

// SolveCodeWars 6
func ReplacePolishLetters(PolishText string) string {
	ReplaceLetters := strings.NewReplacer("ą", "a", "ć", "c", "ę", "e", "ł", "l", "ń", "n", "ó", "o", "ś", "s", "ź", "z", "ż", "z")
	CovertedText := ReplaceLetters.Replace(PolishText)
	return CovertedText
}

// SolveCodeWars 7
func FindAll(OurNumber int, OurArray []int) []int {
	var Array []int
	for index, Number := range OurArray {
		if Number == OurNumber {
			Array = append(Array, index)
		}
	}
	return Array
}

// SolveCodeWars 8
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
