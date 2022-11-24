package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Printf("=======\nЗадача A\n=======\n")
	a, b := 0.4, 0.8
	startValueForX := 3.2
	endValueForX := 6.2
	step := 0.6
	for x := startValueForX; x <= endValueForX; x += step {
		internal.PrintFunctionValue(x, internal.EvaluateFunction(a, b, x))
	}

	fmt.Printf("=======\nЗадача B\n=======\n")
	variableValues := [5]float64{4.48, 3.56, 2.78, 5.28, 3.21}
	for _, x := range variableValues {
		internal.PrintFunctionValue(x, internal.EvaluateFunction(a, b, x))
	}
}
