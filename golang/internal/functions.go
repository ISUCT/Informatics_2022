package internal

import (
	"fmt"
	"math"
)

func EvaluateFunction(a, b, x float64) float64 {
	var functionValue float64 = ((math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a / b))) * math.Pow(a*b, 1.0/3.0)
	return functionValue
}

func PrintFunctionValue(x, functionValue float64) {
	fmt.Printf("f(%.3g) = %.13g\n", x, functionValue)
}
