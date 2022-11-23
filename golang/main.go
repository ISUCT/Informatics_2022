package main

import (
	"fmt"
)

func main() {
	const a float64 = 4.1
	const b float64 = 2.7

	fmt.Println(taskA(a, b))
	fmt.Println(taskB(a, b))
	test()

	fmt.Println("Задания с Codewars")
	fmt.Println(evenOrOdd(16))
	fmt.Println(countingSheep([]bool{true, false, true, true, true}))
	fmt.Println(monkeysCount(19))
	fmt.Println(paperworkCount(23, 2))
	fmt.Println(heroWithGunShootsDragons(37, 20))
}
