package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	const a float64 = 4.1
	const b float64 = 2.7

	fmt.Println(internal.TaskA(a, b))
	fmt.Println(internal.TaskB(a, b))

	fmt.Println("Задания с Codewars")
	fmt.Println(internal.EvenOrOdd(16))
	fmt.Println(internal.CountingSheep([]bool{true, false, true, true, true}))
	fmt.Println(internal.MonkeysCount(19))
	fmt.Println(internal.PaperworkCount(23, 2))
	fmt.Println(internal.HeroWithGunShootsDragons(37, 20))

	//internal.GameOfLife(1, 5, 5)
}
