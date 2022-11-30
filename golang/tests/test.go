package main_test

import (
	"fmt"
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

	fmt.Println("Task A: Test inputs")
	if (xn < xk && xd > 0) || (xn > xk && xd < 0) {
		fmt.Print("- Passed")
	} else {
		fmt.Print("- Failed :", xn, xk, xd)
	}

	fmt.Println("Task A: Test value")
	if assert.InDelta(t, resultA, testResA[1], 0.001) {
		fmt.Print("- Passed")
	} else {
		fmt.Print("- Failed :", testResA[1])
	}
	fmt.Println("Task A: Test length")
	if assert.InDelta(t, 6, len(testResA), 0) {
		fmt.Print("- Passed")
	} else {
		fmt.Print("- Failed :", len(testResA))
	}
}

func TestTaskB(t *testing.T) {
	t.Parallel()

	var resultB float64 = -879771.242

	const a float64 = 4.1
	const b float64 = 2.7

	var testResB []float64 = internal.TaskB(a, b)

	fmt.Println("Task B: Test value")
	if assert.InDelta(t, resultB, testResB[1], 0.001) {
		fmt.Print("- Passed")
	} else {
		fmt.Print("- Failed :", testResB[1])
	}
	fmt.Println("Task B: Test length")
	if assert.InDelta(t, 6, len(testResB), 0) {
		fmt.Print("- Passed")
	} else {
		fmt.Print("- Failed :", len(testResB))
	}
}
