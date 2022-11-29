package main

import "fmt"

func main() {
	for x := 1.2; x <= 4.2; x += 0.6 {
		fmt.Println(lab(2, x))
	}
	fmt.Println(lab(2, 1.16))
	fmt.Println(lab(2, 1.32))
	fmt.Println(lab(2, 1.47))
	fmt.Println(lab(2, 1.65))
	fmt.Println(lab(2, 1.93))
}
