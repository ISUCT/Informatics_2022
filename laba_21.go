package main

import (
	"fmt"
	"math"
)

func main() {
  fmt.Println("A: ")
  for x := 0.11; x <= 0.36; x += 0.05{
	sin := math.Sin(x) 
	cos := math.Cos(x)
	ln := math.Log(x) //натуральный логарифм x  
	fin := (math.Pow(sin, 3) + math.Pow(cos, 3)) * ln //конечная формула
	fmt.Println(fin)
  }
	//var x [5]float64 = [5]float64 {0.2, 0.3, 0.38, 0.43, 0.57}
	//sin := math.Sin(x) 
	//cos := math.Cos(x)
	//ln := math.Log(x)  
	//fin := (math.Pow(sin, 3) + math.Pow(cos, 3)) * ln 
	//fmt.Println(fin)
}
