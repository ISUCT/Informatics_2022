package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println(internal.EvenOrOdd(1))
	fmt.Println(internal.CountingSheeps([]bool{true, false, true, true, true}))
	fmt.Println(internal.CountingMonkeys(10))
	fmt.Println(internal.SchoolPaperwork(-2, 5))
	fmt.Println(internal.IsHe(5, 3))
	fmt.Println(internal.Polish("Jędrzej Błądziński")) //ПОЛЬСКА НЕ РОБИТ
	fmt.Println(internal.FindAll([]int{1, 2, 1, 2, 9, 3, 5, 1}, 1))
	var exampleForSumOfMin [][]int = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	fmt.Println(internal.SumOfMin(exampleForSumOfMin, 3, 5))
}
