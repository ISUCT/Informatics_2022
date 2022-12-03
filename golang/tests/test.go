package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal"
)

func TestTaskA(t *testing.T) {
	t.Parallel()

	var resultA float64 = -2867.204

	const a float64 = 4.1
	const b float64 = 2.7
	const xn float64 = 1.5
	const xk float64 = 3.5
	const xd float64 = 0.4

	var testResA []float64 = internal.TaskA(a, b, xn, xk, xd)

	t.Run("Check 1 value", func(t *testing.T) {
		t.Parallel()
		assert.InDelta(t, resultA, testResA[1], 0.001)
	})
	t.Run("Check xn, xk and xd > 0", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, true, (xn < xk && xd > 0) || (xn > xk && xd < 0))
	})
	t.Run("Check xd compared to xk and xn", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, true, (xd < xk-xn) || (xd > xn-xk))
	})
	t.Run("Check xd compared to xk and xn", func(t *testing.T) {
		t.Parallel()
		assert.InDelta(t, 6, len(testResA), 0)
	})
}

func TestTaskB(t *testing.T) {
	t.Parallel()

	var resultB float64 = -879771.242

	const a float64 = 4.1
	const b float64 = 2.7

	var testResB []float64 = internal.TaskB(a, b, []float64{1.9, 2.15, 2.34, 2.74, 3.16})

	assert.InDelta(t, resultB, testResB[1], 0.001)

	assert.Equal(t, 6, len(testResB))
	assert.Equal(t, []float64{}, internal.TaskB(a, b, []float64{}))
}

func TestFormula(t *testing.T) {
	t.Parallel()

	const a float64 = 4.1
	const b float64 = 2.7

	// Не уверен, что строка ниже вообще будет работать, но пусть будет
	assert.Nil(t, internal.Formula(1, a, b))

	assert.InDelta(t, 3576.30, internal.Formula(3, a, b), 0.01)

}
