package main

import (
	"fmt"
	"math"
)

func main() {
	//task 7
	fmt.Println("TASK A")
	solveTaskA(0.4, 0.8, 3.2, 6.2, 0.6)
	fmt.Println("TASK B")
	args := [5]float64 { 4.48, 3.56, 2.78, 5.28, 3.21 }
	solveTaskB(0.4, 0.8, args)

	//codewars 1
	fmt.Println(evenOddCheck(2))
	fmt.Println(evenOddCheck(5))
	
	//codewars 2
	sheep := []bool {
		true,  true,  true,  false,
		true,  true,  true,  true,
		true,  false, true,  false,
		true,  false, false, true,
		true,  true,  true,  true,
		false, false, true,  true }
	fmt.Println(countSheep(sheep))

	//codewars 3
	fmt.Println(fillArray(10))
	fmt.Println(fillArray(1))

	//codewars 4
	fmt.Println(countPages(5,5))
	fmt.Println(countPages(-5,5))

	//codewars 5
	fmt.Println(isHeroGonnaSurvive(2,1))
	fmt.Println(isHeroGonnaSurvive(5,100))
	fmt.Println(isHeroGonnaSurvive(100,5))
}

//task 7
func solveFunction(a float64, b float64, x float64) float64 {
	return (((math.Pow(a, x) - math.Pow(b, x)) / math.Log10(a / b)) * math.Pow((a * b), 1.0 / 3.0))
}

func solveTaskA(a float64, b float64, xs float64, xe float64, dx float64) {
	for x:=xs; x<=xe; x+=dx {
		fmt.Println(solveFunction(a, b, x))
	} 
}

func solveTaskB(a float64, b float64, args [5]float64) {
	for i:=0; i<len(args); i++ {
		fmt.Println(solveFunction(a, b, args[i]))
	}
}

//codewars 1
func evenOddCheck(num int) string {
	if num % 2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

//codewars 2
func countSheep(sheepArray []bool) int {
	var sheepCount = 0
	for i:=0; i<len(sheepArray); i++ {
		if(sheepArray[i]) {
			sheepCount += 1
		}
	}
	return sheepCount
}

//codewars 3
func fillArray(length int) []int {
	array := []int {}
	for i:=1; i<=length; i++ {
		array = append(array,i)
	}
	return array
}

//codewars 4
func countPages(n int, m int) int {
	if(n<0 || m<0) {
		return 0
	}
	return n * m
}

//codewars 5
func isHeroGonnaSurvive(bulletCount int, dragonCount int) bool {
	var bulletCountToWinDragons = dragonCount * 2
	return bulletCount >= bulletCountToWinDragons
}