package main

import (
	"fmt"
)

func main() {
	//Лабораторная
	for x := 3.5; x <= 6.5; x = x + 0.6 {
		fmt.Println(lab(-2.5, -3.4, x))
	}
	fmt.Println(lab(-2.5, -3.4, 2.89))
	fmt.Println(lab(-2.5, -3.4, 3.54))
	fmt.Println(lab(-2.5, -3.4, 5.21))
	fmt.Println(lab(-2.5, -3.4, 6.28))
	fmt.Println(lab(-2.5, -3.4, 3.48))
	fmt.Println(lab(-2.1, -3.3, 2.85)) //Эта строка сделана только для того, чтобы не ругался линтер
	//А здесь задания с кодварса. Указаны в том порядке, в котором были даны
	fmt.Println(evenorodd(2))
	fmt.Println(sheeps([]bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}))
	fmt.Println(monkeys(10))
	fmt.Println(paper(4, 4))
	fmt.Println(hero(10, 5))
	fmt.Println(polish("Jędrzej Błądziński"))
	fmt.Println(narray([]int{6, 9, 3, 4, 3, 82, 11}, 3))
	fmt.Println(sumofmin([][]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}))
}
