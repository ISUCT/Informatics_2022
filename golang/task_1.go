package main

import "fmt"

func main() {
	var number int
	fmt.Println("Write a number > ")
	fmt.Scanln(&number)
	fmt.Println(even_or_odd(number))
}

func even_or_odd(x int) (answer string) {
	if (x % 2) == 0 {
		answer = "Even"
	} else {
		answer = "Odd"
	}
	return
}
