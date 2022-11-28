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
	fmt.Println(evenOrOdd(16))
	fmt.Println(countingSheep([]bool{true, false, true, true, true}))
	fmt.Println(monkeysCount(19))
	fmt.Println(paperworkCount(23, 2))
	fmt.Println(heroWithGunShootsDragons(37, 20))

	gameOfLife(1, 5, 5)
}
