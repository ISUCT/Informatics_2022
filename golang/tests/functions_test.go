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
	want    float64
}

// Кейс с обычными значениями, должно вернуть значение алгебраической функции
func TestEvaluateFunctionNormal(t *testing.T) {
	testValues := testsEvaluateFunction{a: 0.4, b: 0.8, x: 3.2, want: 0.991503}
	functionValue, err := internal.EvaluateFunction(testValues.a, testValues.b, testValues.x)
	if assert.NoError(t, err) {
		assert.InDelta(t, testValues.want, functionValue, 0.001, fmt.Sprintf("expected: %g got: %g", testValues.want, functionValue))
	}
}

// Проверка на то, возвращает ли функция NaN при заданных значениях
func TestEvaluateFunctionNaN(t *testing.T) {
	testValues := testsEvaluateFunction{a: 0.0, b: 0.0, x: 0.0}
	_, err := internal.EvaluateFunction(testValues.a, testValues.b, testValues.x)
	if assert.Error(t, err) {
		expectedError := errors.New("Not a Number")
		assert.Equal(t, expectedError, err, fmt.Sprintf("expected: %v, got: %v", expectedError, err.Error()))
	}
}

type testsTaskA struct {
	a, b                               float64
	startValueForX, endValueForX, step float64
	want                               int
}

// Кейс с обычными значениями, должно вернуть верную длину слайса
func TestSolveTaskA_Normal(t *testing.T) {
	testValues := testsTaskA{a: 1.2, b: 0.8, startValueForX: 3.2, endValueForX: 6.2, step: 1.2, want: 3}
	slice, err := internal.SolveTaskA(testValues.a, testValues.b, testValues.startValueForX, testValues.endValueForX, testValues.step)
	if assert.NoError(t, err) {
		got := len(slice)
		assert.Equal(t, testValues.want, got, "expected: %d, got: %d", testValues.want, got)
	}
}

// Кейс, где конечное значение x меньше начального значения x,
// должно вернуть ошибку endValueForX is less than startValueForX
func TestSolveTaskA_StartAndEndValues(t *testing.T) {
	testValues := testsTaskA{a: 0.4, b: 0.8, startValueForX: 6.2, endValueForX: 3.2, step: 0.6}
	_, err := internal.SolveTaskA(testValues.a, testValues.b, testValues.startValueForX, testValues.endValueForX, testValues.step)
	if assert.Error(t, err) {
		expectedError := errors.New("endValueForX is less than startValueForX")
		assert.Equal(t, expectedError, err, fmt.Sprintf("expected: %v, got: %v", expectedError, err.Error()))
	}
}

// Кейс, где начальное значение x меньше конечного значения x,
// но шаг отрицательный, должно вернуть ошибку Infinite Cycle
func TestSolveTaskA_NegativeStep(t *testing.T) {
	testValues := testsTaskA{a: 0.8, b: 3.2, startValueForX: 3.2, endValueForX: 6.2, step: -0.6}
	_, err := internal.SolveTaskA(testValues.a, testValues.b, testValues.startValueForX, testValues.endValueForX, testValues.step)
	if assert.Error(t, err) {
		expectedError := errors.New("Infinite cycle")
		assert.Equal(t, expectedError, err, fmt.Sprintf("expected: %v, got: %v", expectedError, err.Error()))
	}
}

type testsTaskB struct {
	a, b       float64
	size       int
	valuesForX []float64
	want       int
}

// Проверка на возвращаемую длину слайса, должно вернуть верную длину слайса
func TestSolveTaskB_ReturnedLength(t *testing.T) {
	testValues := testsTaskB{a: 0.4, b: 0.8, size: 5, valuesForX: []float64{4.48, 3.56, 2.78, 5.28, 3.21, 7.0, 10.0}, want: 5}
	got := len(internal.SolveTaskB(testValues.a, testValues.b, testValues.size, testValues.valuesForX))
	assert.Equal(t, testValues.want, got, "expected: %d, got: %d", testValues.want, got)
}
