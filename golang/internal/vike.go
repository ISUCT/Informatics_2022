package internal

import (
	"fmt"
	"math"
)

func ReshX(a float64, c float64) float64 {
	return (((math.Pow(math.Log(math.Abs(3*a+c)), 2)) / 5) + math.Sin(math.Pi*c))
}
func ReshY(a float64, b float64) float64 {
	return (math.Abs(1-a) * math.Sin((math.Pi/4)+b))
}
func Vike() {
	fmt.Println(ReshX(1, 2))
	fmt.Println(ReshY(1, 2))
}
