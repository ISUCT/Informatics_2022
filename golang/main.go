package main

import (
	"flag"
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	var a, b float64
	var startValueForX, endValueForX, step float64
	var includeCustomSlice bool

	flag.Float64Var(&a, "a", 0.4, "sets value of variable a")
	flag.Float64Var(&b, "b", 0.8, "sets value of variable b")
	flag.Float64Var(&startValueForX, "start", 3.2, "sets the start x value")
	flag.Float64Var(&endValueForX, "end", 6.2, "sets the end x value")
	flag.Float64Var(&step, "step", 0.6, "sets the delta x")
	flag.BoolVar(&includeCustomSlice, "include-slice", false, "set this flag to true to include user-inputted slice")

	flag.Parse()

	fmt.Println("Task A:")
	answerTaskA, _ := internal.SolveTaskA(a, b, startValueForX, endValueForX, step)
	internal.PrintFunctionValue(answerTaskA)

	fmt.Println("Task B:")
	var variableValues = []float64{4.48, 3.56, 2.78, 5.28, 3.21}

	if includeCustomSlice {
		variableValues = internal.ReadUserInput()
	}
	answerTaskB := internal.SolveTaskB(a, b, variableValues)
	internal.PrintFunctionValue(answerTaskB)
}
