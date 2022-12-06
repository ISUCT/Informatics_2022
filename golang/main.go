package main

import "fmt"

//HW Chukhlova V

func EvenOrOdd(num int) {
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func Sheeps(num []bool) {
	var count = 0
	for _, value := range num {
		if value == true {
			count++
		}
	}
	fmt.Println(count)
}

func Paperworks(n int, m int) {
	if n > 0 && m > 0 {
		fmt.Println(n * m)
	} else {
		fmt.Println(0)
	}
}

func Dragons(bullets int, dragons int) {
	if bullets/dragons >= 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
}
