package main

import (
	"fmt"
	"strconv"

	"isuct.ru/informatics2022/internal"
)

func main() {
	const a float64 = 4.1
	const b float64 = 2.7
	const xn float64 = 1.5
	const xk float64 = 3.5
	const xd float64 = 0.4

	fmt.Println(internal.TaskA(a, b, xn, xk, xd))
	fmt.Println(internal.TaskB(a, b, []float64{1.9, 2.15, 2.34, 2.74, 3.16}))

	fmt.Println("Задания с Codewars")
	fmt.Println(internal.EvenOrOdd(16))
	fmt.Println(internal.CountingSheep([]bool{true, false, true, true, true}))
	fmt.Println(internal.MonkeysCount(19))
	fmt.Println(internal.PaperworkCount(23, 2))
	fmt.Println(internal.HeroWithGunShootsDragons(37, 20))

	sample_grid := [][]byte{
		{2, 2, 2, 2, 2, 2, 2},
		{2, 0, 0, 0, 0, 0, 2},
		{2, 0, 0, 1, 0, 0, 2},
		{2, 1, 0, 1, 0, 0, 2},
		{2, 0, 1, 1, 0, 0, 2},
		{2, 0, 0, 0, 0, 0, 2},
		{2, 2, 2, 2, 2, 2, 2},
	}

	fmt.Println(internal.PolishCow("Gdzie jest biały węgorz?"))

	//Code to convert []int to []string
	nums := []int{6, 9, 3, 4, 3, 82, 11}
	sNums := make([]string, len(nums))
	for i, x := range nums {
		sNums[i] = strconv.Itoa(x)
	}

	fmt.Println(internal.FindAll(sNums, fmt.Sprint(3)))

	var exampleArr = [][]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
		{20, 21, 34, 56, 100},
	}

	fmt.Println(internal.SumOfMin(exampleArr, 3, 5))

	internal.GameOfLife(5, 5, 5, sample_grid)
}
