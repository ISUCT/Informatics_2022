//Хан Анна 1/147 Вариант 22

package main

import "fmt"
import "math"


func main() {
	fmt.Println("Hello world")
	var a = 2.25
	fmt.Println("Task А")
	for x := 1.2; x <= 2.7; x += 0.3 {
		y := (math.Pow(a, (math.Pow(x,2.0))-1.0)-math.Log10((math.Pow(x,2.0))-1.0)+math.Pow((math.Pow(x,2.0))-1.0,3.0/2.0))
		fmt.Println("x = ", x, " ", "y = ", y)
	}
	var mass = [5]float64{1.31, 1.39, 1.44, 1.56, 1.92}
	fmt.Println("Task B")
	for _, x1 := range mass {
		y1 := (math.Pow(a, math.Pow(x1,2.0)-1.0)-math.Log10((math.Pow(x1,2.0))-1.0)+math.Pow((math.Pow(x1,2.0))-1.0,3.0/2.0))
		fmt.Println("x = ", x1, " ", "y = ", y1)
	}
}
