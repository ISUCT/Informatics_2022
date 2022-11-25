package internal_test

import (
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
		assert.Equal(fmt.Errorf("Not a Number"), err, "Should return Not a Number.")
	}

	a, b = -1.0, -5.0
	x = -3.5
	functionValue, err = internal.EvaluateFunction(a, b, x)
	// Pow(x, y) = NaN for finite x < 0 and finite non-integer y
	if assert.Error(err) {
		assert.Equal(fmt.Errorf("Not a Number"), err, "Should return Not a Number.")
	}
}

func TestSolveTaskA(t *testing.T) {
	assert := assert.New(t)

	var a, b float64 = 0.4, 0.8
	var want int = 6
	var got int = len(internal.SolveTaskA(a, b))

	assert.Equal(want, got, "Want and got should be equal.")
}

func TestSolveTaskB(t *testing.T) {
	assert := assert.New(t)

	var a, b float64 = 0.4, 0.8
	var want int = 5
	var got int = len(internal.SolveTaskB(a, b))

	assert.Equal(want, got, "Want and got should be equal.")
}
