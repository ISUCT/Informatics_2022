package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	const a float64 = 4.1
	const b float64 = 2.7
	const xn float64 = 1.5
	const xk float64 = 3.5
	const xd float64 = 0.4

	fmt.Println(internal.TaskA(a, b, xn, xk, xd))
	fmt.Println(internal.TaskB(a, b))

	fmt.Println("Задания с Codewars")
	fmt.Println(internal.EvenOrOdd(16))
	fmt.Println(internal.CountingSheep([]bool{true, false, true, true, true}))
	fmt.Println(internal.MonkeysCount(19))
	fmt.Println(internal.PaperworkCount(23, 2))
	fmt.Println(internal.HeroWithGunShootsDragons(37, 20))

	//internal.GameOfLife(1, 5, 5)
}
