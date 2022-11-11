package main

import (
	"fmt"
	"math"
)

func f (x float64) float64 {
  y := (math.Pow(math.Sin(x), 3) + math.Pow(math.Cos(x), 3)) * math.Log(x)
  return y
}

//Задание А:
func main() {
  fmt.Println("A: ")
  for x := 0.11; x <= 0.36; x += 0.05 {
    sin := math.Sin(x) 
    cos := math.Cos(x)
	  ln := math.Log(x) //натуральный логарифм x  
	  fin := (math.Pow(sin, 3) + math.Pow(cos, 3)) * ln //конечная формула
    fmt.Println(fin)
	}
  fmt.Println("B: ")
  fmt.Println(f(0.2))
  fmt.Println(f(0.3))
  fmt.Println(f(0.38))
  fmt.Println(f(0.43))
  fmt.Println(f(0.57))
} 
