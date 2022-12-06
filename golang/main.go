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
  fmt.Println(monkeyCount(3))
} 
//Codewars 
//1
func EvenOrOdd(number int) string {
    if number%2 == 0{
        return "Even"
    } else {
        return "Odd"
    }
}
//2 
func CountSheeps(numbers []bool) int {
  var count = 0
  for i:=0; i<len(numbers); i++ {
    if(numbers[i]) {
      count += 1
      }
    }
    return count
}
//3
func monkeyCount(n int) []int {
  var r = []int{}
  for i:= 1; i <= n; i++{
    r = append (r, i)
  }
  return r
}
//4 - невозможно решить, т.к отсутствует вариант этой задачи для Golang
//5
func Hero(bullets, dragons int) bool {
  return bullets >= dragons * 2
}
