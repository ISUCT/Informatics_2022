package main

import "fmt"

func main() {

}

func evenorodd(s int) {
	fmt.Println("1st task from codewars")
	if s%2 == 0 {
		fmt.Println("четное")
	} else {
		fmt.Println("нечетное")
	}
}

func countingsheeps(numbers []bool) int {
	fmt.Println("2nd task from codewars")
	var count = 0

	for _, i := range numbers {
		// просчитывание массива, используем пустую переменную, чтобы не было ошибки but not used
		if i {
			count += 1
		}

	}
	return (count)
}

func monkey(n int) []int {
	fmt.Println("3d task from codewars")
	monc := make([]int, n)
	// команда make позволяет вывести массив из n переменных типа int с пустым значением
	for i := range monc {
		//диапозон monc
		monc[i] = i + 1

	}

	return (monc)
}

func classmates(n, m int) {
	fmt.Println("4th task from codewars")
	if n < 0 || m < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(n * m)
	}
	//если я буду использовать return будет ошибка too many return values

}

//у нас есть какое-то кол-во драконов x, и какое-то количество пуль y. если мы возьмем у/2<x, то умрем

func dragons(x, y int) {
	fmt.Println("5th task from codewars")
	if y/2 < x {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}
