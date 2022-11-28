package internal

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func EvaluateFunction(a, b, x float64) (float64, error) {
	var functionValue float64 = ((math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a / b))) * math.Cbrt(a*b)
	if functionValue != functionValue { // NaN != NaN (always)
		return 0, errors.New("Not a Number")
	}
	return functionValue, nil
}

func PrintFunctionValue(functionValue []float64) {
	for _, elem := range functionValue {
		fmt.Println(elem)
	}
}

func SolveTaskA(a, b, startValueForX, endValueForX, step float64) ([]float64, error) {
	if endValueForX < startValueForX {
		return nil, errors.New("endValueForX is less than startValueForX")
	}
	if (endValueForX > startValueForX) && (step < 0) {
		return nil, errors.New("Infinite cycle")
	}

	size := math.Floor(math.Abs((endValueForX-startValueForX)/step)) + 1 // количество членов арифметической прогрессии
	if size <= 0 {
		log.Fatal("Size is less or equal to zero")
	}
	sliceOfResults := make([]float64, 0, int(size))

	for x := startValueForX; x <= endValueForX; x += step {
		functionValue, _ := EvaluateFunction(a, b, x)
		sliceOfResults = append(sliceOfResults, functionValue)
	}
	return sliceOfResults, nil
}

func SolveTaskB(a, b float64, size int, variableValues []float64) []float64 {
	sliceOfResults := make([]float64, 0, size)

	for i := 0; i < size; i++ {
		functionValue, err := EvaluateFunction(a, b, variableValues[i])
		if err != nil {
			log.Fatal(err.Error())
		}
		sliceOfResults = append(sliceOfResults, functionValue)
	}
	return sliceOfResults
}
