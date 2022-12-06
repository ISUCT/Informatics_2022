package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Type n > ")
	fmt.Scanln(&n)
	fmt.Print("Type m > ")
	fmt.Scanln(&m)
	if (n < 0) || (m < 0) {
		fmt.Println(0)
	} else {
		fmt.Println(n * m)
	}
}
