/*
package main

import (

	"fmt"
	"math"

)

	func formula(x float64) float64 {
		var a float64 = 4.1
		var b float64 = 2.7
		var y float64 = (a*(math.Cbrt(x)) - (b * (math.Log(x) / math.Log(5)))) / (math.Log10(math.Abs(x - 1)))
		return y
	}

	func main() {
		fmt.Println("Задание А")
		var i float64 = 1.2
		for ; i < 5.2; i = i + 0.8 {
			fmt.Println(formula(i))
		}
		fmt.Println("Задание B")
		fmt.Println(formula(1.9))
		fmt.Println(formula(2.15))
		fmt.Println(formula(2.34))
		fmt.Println(formula(2.73))
		fmt.Println(formula(3.16))
	}
*/
package main

import (
	"fmt"
	"math"

	"isuct.ru/informatics2022/internal"
)

func Formula(x float64, a float64, b float64) float64 {
	var y float64 = (a*(math.Cbrt(x)) - (b * (math.Log(x) / math.Log(5)))) / (math.Log10(math.Abs(x - 1)))
	return y
}

func TaskA(a float64, b float64, xn float64, xk float64, xd float64) []float64 {
	var res []float64
	for i := xn; i < xk; i = i + xd {
		res = append(res, Formula(i, a, b))
	}
	return res
}
func TaskB(a float64, b float64, x []float64) []float64 {
	var res []float64
	for i := 0; i < len(x); i++ {
		res = append(res, Formula(x[i], a, b))
	}
	return res
}
func main() {
	const a float64 = 4.1
	const b float64 = 2.7
	const xn float64 = 1.2
	const xk float64 = 5.2
	const xd float64 = 0.8

	fmt.Println(TaskA(a, b, xn, xk, xd))
	fmt.Println(TaskB(a, b, []float64{1.9, 2.15, 2.34, 2.73, 3.16}))
	fmt.Println(internal.EvenOdd(666))
}
