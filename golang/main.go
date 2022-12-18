package main

import (
	"fmt"
	"math"
)

// lab 3 var 7
func main() {
	fmt.Println("Hello world")
	fmt.Println("Num A")
	numAResults := numA(0.4, 0.8, 3.2, 6.2, 0.6)
	for i := 0; i < len(numAResults); i++ {
		fmt.Println(numAResults[i])
	}
	fmt.Println("Mass")
	digit := []float64{4.48, 3.56, 2.78, 5.28, 3.21}
	Mass(0.4, 0.8, digit)
	fmt.Println("Codewars tasks ") //Codewars tasks
	// codewars 1
	fmt.Println(cod1(3), cod1(4))
	// codewars 2
	count := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}
	fmt.Println(cod2(count))
	// codewars 3
	fmt.Println(cod3(10), cod3(1))

	// codewars 4
	fmt.Println(cod4(5, 5), cod4(0, 5))
	// codewars 5
	fmt.Println(cod5(20, 5))

}
func lab(a float64, b float64, x float64) float64 {
	return (((math.Pow(a, x) - math.Pow(b, x)) / math.Log10(a/b)) * math.Pow((a*b), 1.0/3.0))
}
func numA(a float64, b float64, xn float64, xk float64, xd float64) []float64 {
	var array []float64
	for x := xn; x < xk; x += xd {

		array = append(array, lab(a, b, x))
	}
	return array
}
func Mass(a float64, b float64, digit []float64) {
	for i := 0; i < len(digit); i++ {
		fmt.Println(lab(a, b, digit[i]))
	}

}

// task  1 codewars
func cod1(n int) string {
	if n%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// task 2 codewars
func cod2(sheep []bool) int {

	var k = 0
	for i := 0; i < len(sheep); i++ {
		if sheep[i] == true {
			k += 1
		}
	}
	return k
}

// codewars 3
func cod3(monkeys int) []int {
	var mass []int
	for i := 1; i <= monkeys; i++ {
		mass = append(mass, i)
	}
	return mass

}

// codewars 4
func cod4(classmates int, paperwork int) int {
	var product int

	if classmates <= 0 {
		product = 0
	} else if classmates > 0 {
		product = classmates * paperwork
	}
	return product
}

// codewars 5
func cod5(bullets int, dragons int) bool {
	var product = dragons * 2
	return bullets >= product

}
