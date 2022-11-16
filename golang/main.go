package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Println(lab(-2.5, 3.4, 5.1))
	for x := 3.5; x <= 6.5; x += 0.6 {
		fmt.Println(lab(-2.5, 3.4, x))
	}
	fmt.Println(lab(-2.5, 3.4, 2.89))
	fmt.Println(lab(-2.5, 3.4, 3.54))
	fmt.Println(lab(-2.5, 3.4, 5.21))
	fmt.Println(lab(-2.5, 3.4, 6.28))
	fmt.Println(lab(-2.5, 3.4, 3.48))
}
