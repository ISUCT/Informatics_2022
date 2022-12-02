package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//Лабораторная работа
	fmt.Println("Task А")
	SolveTaskA(4.1, 2.7, 1.2, 5.2, 0.8)

	fmt.Println("Task B")
	MassTaskB := [5]float64{1.9, 2.15, 2.34, 2.73, 3.16}
	SolveTaskB(4.1, 2.7, MassTaskB)

	//CodeWars 1
	fmt.Println("CodeWarsTask 1")
	fmt.Println(CheckOddOrEven(777), ",", CheckOddOrEven(666))

	//CodeWars 2
	fmt.Println("CodeWarsTask 2")
	Sheep := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	fmt.Println(SheepsCounter(Sheep))

	//CodeWars 3
	fmt.Println("CodeWarsTask 3")
	fmt.Println(MonkeysCounter(10))
	fmt.Println(MonkeysCounter(1))

	//CodeWars 4
	fmt.Println("CodeWarsTask 4")
	fmt.Println(CountPaperwork(5, 5))
	fmt.Println(CountPaperwork(-5, 5))

	//CodeWars 5
	fmt.Println("CodeWarsTask 5")
	fmt.Println(IsHeGonnaSurvive(5, 2))
	fmt.Println(IsHeGonnaSurvive(1, 2))

	//CodeWars 6
	fmt.Println("CodeWarsTask 6")
	fmt.Println(ReplacePolishLetters("Jędrzej Błądziński"))

	//CodeWars 7
	fmt.Println("CodeWarsTask 7")
	FindAllExampleArray := []int{6, 9, 3, 4, 3, 82, 11}
	fmt.Println(FindAll(3, FindAllExampleArray))
	fmt.Println(FindAll(1, FindAllExampleArray))

	//CodeWars 8
	fmt.Println("CodeWarsTask 8")
	SumOfMinimumsExampleArray := [][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 34, 56, 100}}
	fmt.Println(SumOfMinimums(SumOfMinimumsExampleArray))

}

// Function For The Formula
func formula(a float64, b float64, x float64) float64 {
	return ((a * math.Sqrt(x)) - (b * (math.Log2(x) / math.Log2(5)))) / (math.Log(math.Abs(x - 1)))
}

// Funtion For The Task A
func SolveTaskA(a float64, b float64, xn float64, xk float64, dx float64) {
	for x := xn; x <= xk; x += dx {
		fmt.Println(formula(a, b, x))
	}
}

// Funtion For The Task B
func SolveTaskB(a float64, b float64, MassTaskB [5]float64) {
	for i := 0; i < len(MassTaskB); i += 1 {
		fmt.Println(formula(a, b, MassTaskB[i]))
	}
}

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
