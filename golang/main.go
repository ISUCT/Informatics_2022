package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	// var a = 2.25
	// fmt.Println("№22")
	// fmt.Println("А")
	// for x := 1.2; x <= 2.7; x += 0.3 {
	// 	var first = math.Pow(a, (math.Pow(x, 2) - 1))
	// 	var second = math.Log(math.Pow(x, 2) - 1)
	// 	var third = math.Pow(math.Pow(x, 2)-1, 1/3)
	// 	var result = first - second + third
	// 	fmt.Println(result)
	// }

	// fmt.Println("B")
	// arr := []float64{1.31, 1.39, 1.44, 1.56, 1.92}
	// for _, b := range arr {
	// 	var first = math.Pow(a, (math.Pow(b, 2) - 1))
	// 	var second = math.Log(math.Pow(b, 2) - 1)
	// 	var third = math.Pow(math.Pow(b, 2)-1, 1/3)
	// 	var result = first - second + third
	// 	fmt.Println(result)
	// }

	fmt.Println("Even or Odd №1:", internal.EvenOrOdd(10))

	fmt.Println("Counting sheep №2:", internal.MonkeyCount(10))

	fmt.Println("monkeyCount №3:", internal.CountSheeps([]bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}))

	fmt.Println("School Paperwork №4:", internal.Paperwork(5, 5))

	fmt.Println("Is he gonna survive? #5:", internal.Hero(10, 5))

	fmt.Println("CorrectPolishLetters #6:", internal.CorrectPolishLetters("Jędrzej Błądziński"))

	fmt.Println("Find all occurrences.. #7:", internal.FindAll(4, []int{6, 9, 3, 4, 3, 82, 4}))

	fmt.Println("Find all occurrences.. #8:", internal.SumOfMinimums([][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 34, 56, 100}}))
}
