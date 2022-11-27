package internal_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"isuct.ru/informatics2022/internal"
)

type testsEvaluateFunction struct {
	a, b, x float64
	want    string
}

func TestEvaluateFunction(t *testing.T) {
	assert := assert.New(t)

	testValues := testsEvaluateFunction{a: 0.4, b: 0.8, x: 3.2, want: "0.991503"}

	functionValue, err := internal.EvaluateFunction(testValues.a, testValues.b, testValues.x)
	assert.Equal(testValues.want, fmt.Sprintf("%.6f", functionValue), "Want and got should be equal.")

	testValues = testsEvaluateFunction{a: 0.0, b: 0.0, x: 0.0}
	functionValue, err = internal.EvaluateFunction(testValues.a, testValues.b, testValues.x)
	if assert.Error(err) {
		assert.Equal(errors.New("Not a Number"), err, "Should return Not a Number.")
	}

	testValues = testsEvaluateFunction{a: -1.0, b: -5.0, x: -3.5}
	functionValue, err = internal.EvaluateFunction(testValues.a, testValues.b, testValues.x)
	if assert.Error(err) { // Pow(x, y) = NaN for finite x < 0 and finite non-integer y
		assert.Equal(errors.New("Not a Number"), err, "Should return Not a Number.")
	}
}

type testsTaskA struct {
	a, b                               float64
	startValueForX, endValueForX, step float64
	want                               int
}

func TestSolveTaskA(t *testing.T) {
	assert := assert.New(t)

	testValues := testsTaskA{a: 0.4, b: 0.8, startValueForX: 3.2, endValueForX: 6.2, step: 0.6, want: 6}
	slice, err := internal.SolveTaskA(testValues.a, testValues.b, testValues.startValueForX, testValues.endValueForX, testValues.step)
	got := len(slice)
	assert.Equal(testValues.want, got, "A length of the resulting array should be equal to 6.")

	testValues = testsTaskA{a: 0.4, b: 0.8, startValueForX: 3.2, endValueForX: 6.2, step: 1.2, want: 3}
	slice, err = internal.SolveTaskA(testValues.a, testValues.b, testValues.startValueForX, testValues.endValueForX, testValues.step)
	got = len(slice)
	assert.Equal(testValues.want, got, "A length of the resulting array should be equal to 3.")

	testValues = testsTaskA{a: 0.4, b: 0.8, startValueForX: 6.2, endValueForX: 3.2, step: 0.6}
	slice, err = internal.SolveTaskA(testValues.a, testValues.b, testValues.startValueForX, testValues.endValueForX, testValues.step)
	if assert.Error(err) {
		assert.Equal(errors.New("endValueForX is less than startValueForX"), err, err.Error())
	}

	testValues = testsTaskA{a: 0.4, b: 0.8, startValueForX: 3.2, endValueForX: 6.2, step: -0.6}
	slice, err = internal.SolveTaskA(testValues.a, testValues.b, testValues.startValueForX, testValues.endValueForX, testValues.step)
	if assert.Error(err) {
		assert.Equal(errors.New("Infinite cycle"), err, err.Error())
	}
}

type testsTaskB struct {
	a, b       float64
	size       int
	valuesForX []float64
	want       int
}

func TestSolveTaskB(t *testing.T) {
	assert := assert.New(t)

	testValues := testsTaskB{a: 0.4, b: 0.8, size: 5, valuesForX: []float64{4.48, 3.56, 2.78, 5.28, 3.21, 7.0, 10.0}, want: 5}

	var got int = len(internal.SolveTaskB(testValues.a, testValues.b, testValues.size, &testValues.valuesForX))

	assert.Equal(testValues.want, got, "A length of the resulting array should be equal to 5.")
}
