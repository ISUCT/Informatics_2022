package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	// Указатели сделаны, чтобы передавать в функции слайс напрямую, а не делать каждый раз его копию
	fmt.Printf("=======\nЗадача A\n=======\n")
	a, b := 0.4, 0.8
	answerTaskA := internal.SolveTaskA(a, b)
	internal.PrintFunctionValue(&answerTaskA)

	fmt.Printf("=======\nЗадача B\n=======\n")
	answerTaskB := internal.SolveTaskB(a, b)
	internal.PrintFunctionValue(&answerTaskB)
}
