package main

import (
	"fmt"
	"math"
)

func EvaluateFunction(a, b, x float64) float64 {
	var functionValue float64 = ((math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a / b))) * math.Pow(a*b, 1.0/3.0)
	return functionValue
}

func printFunctionValue(x, functionValue float64) {
	fmt.Printf("f(%.3g) = %.13g\n", x, functionValue)
}

func main() {
	fmt.Printf("=======\nЗадача A\n=======\n")
	a, b := 0.4, 0.8
	startValueForX := 3.2
	endValueForX := 6.2
	step := 0.6
	for x := startValueForX; x <= endValueForX; x += step {
		printFunctionValue(x, EvaluateFunction(a, b, x))
	}

	fmt.Printf("=======\nЗадача B\n=======\n")
	variableValues := [5]float64{4.48, 3.56, 2.78, 5.28, 3.21}
	for _, x := range variableValues {
		printFunctionValue(x, EvaluateFunction(a, b, x))
	}
}
