package main

import "fmt"

func test() {
	var resultsA = [6]float64{-2867.2049337558983, -879771.242424423, 60553.396028559044, 7705.642156268663, 2947.1559153687504, 1627.2100029870098}
	var resultsB = [5]float64{-879771.242424423, 392039.14018305, 43864.686778786934, 6807.290024311177, 2651.777897551827}

	const a float64 = 4.1
	const b float64 = 2.7

	var testResA []float64 = taskA(a, b)
	var testResB []float64 = taskB(a, b)

	for i := 0; i < len(resultsA); i++ {
		if resultsA[i] == testResA[i] {
			fmt.Println("Task A: Test ", i+1, "- Passed")
		} else {
			fmt.Println("Task A: Test ", i+1, "- Failed :", testResA[i])
		}
	}
	fmt.Println("-----------")
	for i := 0; i < len(resultsB); i++ {
		if resultsB[i] == testResB[i] {
			fmt.Println("Task B: Test ", i+1, "- Passed")
		} else {
			fmt.Println("Task B: Test ", i+1, "- Failed", testResB[i])
		}
	}
}
