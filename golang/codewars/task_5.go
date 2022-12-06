package main

import "fmt"

func main() {
	fmt.Print("How many bullets hero has? > ")
	var bullets int
	fmt.Scanln(&bullets)
	fmt.Print("How many dragons awaiting? > ")
	var dragons int
	fmt.Scanln(&dragons)
	if float32(bullets/2.0) >= float32(dragons) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
