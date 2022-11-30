package main

import (
	"flag"
	"fmt"
	"log"

	"isuct.ru/informatics2022/internal"
)

func main() {
	useDefaultValues := flag.Bool("d", true, "Omit this flag to use default values, overwise set this flag to false and type your own values.")
	flag.Parse()

	if *useDefaultValues {
		fmt.Println("Task A:")
		answerTaskA, err := internal.SolveTaskA(0.4, 0.8, 3.2, 6.2, 0.6)
		if err != nil {
			log.Fatal(err)
		}
		internal.PrintFunctionValue(answerTaskA)

		fmt.Println("Task B:")
		var variableValues = []float64{4.48, 3.56, 2.78, 5.28, 3.21}
		answerTaskB := internal.SolveTaskB(0.4, 0.8, variableValues)
		internal.PrintFunctionValue(answerTaskB)

		return
	}

	fmt.Println("Task A: please, input a, b, start value for x, end value for x and step for x")
	var a, b float64
	var startValueForX, endValueForX, step float64
	fmt.Println("Task A: please, input a, b, start value for x, end value for x and step for x")
	_, err := fmt.Scan(&a, &b, &startValueForX, &endValueForX, &step)
	if err != nil {
		log.Fatal(err)
	}

	answerTaskA, err := internal.SolveTaskA(a, b, startValueForX, endValueForX, step)
	if err != nil {
		log.Fatal(err)
	}
	internal.PrintFunctionValue(answerTaskA)

	fmt.Println("Task B: please, enter a, b parameters and values for x itself")
	_, err = fmt.Scan(&a, &b)
	if err != nil {
		log.Fatal(err)
	}

	variableValues := internal.ReadUserInput()
	answerTaskB := internal.SolveTaskB(a, b, variableValues)
	internal.PrintFunctionValue(answerTaskB)
}
