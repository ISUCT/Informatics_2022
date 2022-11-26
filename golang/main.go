package main

import (
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Printf("=======\nЗадача A\n=======\n")
	var a, b float64
	var startValueForX, endValueForX, step float64
	fmt.Println("Task A: please, input a, b, start value for x, end value for x and step for x") // Указатели сделаны, чтобы передавать в функции слайс
	_, err := fmt.Scan(&a, &b, &startValueForX, &endValueForX, &step)                            // напрямую, а не делать каждый раз его копию
	if err != nil {
		log.Fatal(err)
	}

	answerTaskA, err := internal.SolveTaskA(a, b, startValueForX, endValueForX, step)
	if err != nil {
		log.Fatal(err)
	}
	internal.PrintFunctionValue(&answerTaskA)

	fmt.Printf("=======\nЗадача B\n=======\n")
	fmt.Println("Task B: please, enter a, b, quantity of values of x and enter values for x themselves")

	var size int
	_, err = fmt.Scan(&a, &b, &size)
	if err != nil {
		log.Fatal(err)
	}

	variableValues := make([]float64, 0, size)
	var variableX float64
	for i := 0; i < size; i++ { // Если введёных элементов больше, чем их указанное количество (size), то "лишние" элементы не вносятся в слайс variableX
		_, err = fmt.Scan(&variableX)
		if err != nil {
			log.Fatal(err)
		}
		variableValues = append(variableValues, variableX)
	}

	answerTaskB := internal.SolveTaskB(a, b, size, &variableValues)
	internal.PrintFunctionValue(&answerTaskB)
}
