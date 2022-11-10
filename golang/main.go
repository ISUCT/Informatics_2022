package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.7; x <= 2.2; x += 0.3 {
		a := 1.2
		b := 0.48
		const e float64 = 2.7182
		var y float64 = (math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos(x*b*x) + (math.Pow(e, -x/2)))
		fmt.Println(y)
	}

}
