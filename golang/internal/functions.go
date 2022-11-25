package internal

import (
	"fmt"
	"math"
)

func EvaluateFunction(a, b, x float64) (float64, error) {
	var functionValue float64 = ((math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a / b))) * math.Cbrt(a*b)
	if functionValue != functionValue { // NaN != NaN (always)
		return 0.0, fmt.Errorf("Not a Number")
	}
	return functionValue, nil
}

func PrintFunctionValue(functionValue *[]float64) {
	for _, elem := range *functionValue {
		fmt.Println(elem)
	}
}

func SolveTaskA(a, b float64) []float64 {
	sliceOfResults := make([]float64, 0, 6)

	startValueForX, endValueForX := 3.2, 6.2
	step := 0.6
	for x := startValueForX; x <= endValueForX; x += step {
		functionValue, _ := EvaluateFunction(a, b, x)
		sliceOfResults = append(sliceOfResults, functionValue)
	}
	return sliceOfResults
}

func SolveTaskB(a, b float64) []float64 {
	sliceOfResults := make([]float64, 0, 5)

	variableValues := [5]float64{4.48, 3.56, 2.78, 5.28, 3.21}
	for _, x := range variableValues {
		functionValue, _ := EvaluateFunction(a, b, x)
		sliceOfResults = append(sliceOfResults, functionValue)
	}
	return sliceOfResults
}
