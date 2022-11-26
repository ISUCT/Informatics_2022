package internal_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"isuct.ru/informatics2022/internal"
)

func TestEvaluateFunction(t *testing.T) {
	assert := assert.New(t)

	var a, b float64 = 0.4, 0.8
	var x float64 = 3.2
	var want string = "0.991503"
	functionValue, err := internal.EvaluateFunction(a, b, x)
	assert.Equal(want, fmt.Sprintf("%.6f", functionValue), "Want and got should be equal.")

	a, b = 0.0, 0.0
	x = 0.0
	functionValue, err = internal.EvaluateFunction(a, b, x)
	if assert.Error(err) {
		assert.Equal(err.Error(), "Not a Number", "Should return Not a Number.")
	}

	a, b = -1.0, -5.0
	x = -3.5
	functionValue, err = internal.EvaluateFunction(a, b, x)
	// Pow(x, y) = NaN for finite x < 0 and finite non-integer y
	if assert.Error(err) {
		assert.Equal(err.Error(), "Not a Number", "Should return Not a Number.")
	}
}

func TestSolveTaskA(t *testing.T) {
	assert := assert.New(t)

	var a, b float64 = 0.4, 0.8
	var startValueForX, endValueForX, step = 3.2, 6.2, 0.6
	var want int = 6
	slice, err := internal.SolveTaskA(a, b, startValueForX, endValueForX, step)
	//if assert.Error(err) {
	//	assert.Equal(nil, err, err.Error())
	//}
	got := len(slice)
	assert.Equal(want, got, "A length of the resulting array should be equal to 6.")

	a, b = 0.4, 0.8
	startValueForX, endValueForX, step = 6.2, 3.2, 0.6
	slice, err = internal.SolveTaskA(a, b, startValueForX, endValueForX, step)
	if assert.Error(err) {
		assert.Equal(errors.New("EndValue is less than StartValue"), err, err.Error())
	}

	a, b = 0.4, 0.8
	startValueForX, endValueForX, step = 3.2, 6.2, -0.6
	slice, err = internal.SolveTaskA(a, b, startValueForX, endValueForX, step)
	if assert.Error(err) {
		assert.Equal(errors.New("Infinite cycle"), err, err.Error())
	}
}

func TestSolveTaskB(t *testing.T) {
	assert := assert.New(t)

	var a, b float64 = 0.4, 0.8
	var size int = 5
	var valuesForX = []float64{4.48, 3.56, 2.78, 5.28, 3.21, 7.0, 10.0}
	var want int = 5
	var got int = len(internal.SolveTaskB(a, b, size, &valuesForX))

	assert.Equal(want, got, "A length of the resulting array should be equal to 5.")
}
