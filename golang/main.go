package main

import (
	"fmt"
	"math"
)

func main() {
	//task 7
	solveTask()
}

func solveFunction(a float64, b float64, x float64) float64 {
	return (((math.Pow(a, x) - math.Pow(b, x)) / math.Log10(a / b)) * math.Pow((a * b), 1.0 / 3.0))
}

func solveTaskA(a float64, b float64, xs float64, xe float64, dx float64) {
	for x:=xs; x<=xe; x+=dx {
		fmt.Println(solveFunction(a, b, x))
	} 
}

func solveTaskB(a float64, b float64, x1 float64, x2 float64, x3 float64, x4 float64, x5 float64) {
	fmt.Println(solveFunction(a, b, x1))
	fmt.Println(solveFunction(a, b, x2))
	fmt.Println(solveFunction(a, b, x3))
	fmt.Println(solveFunction(a, b, x4))
	fmt.Println(solveFunction(a, b, x5))
}

func solveTask() {
	fmt.Println("TASK A")
	solveTaskA(0.4, 0.8, 3.2, 6.2, 0.6)
	fmt.Println("TASK B")
	solveTaskB(0.4, 0.8, 4.48, 3.56, 2.78, 5.28, 3.21)
}
