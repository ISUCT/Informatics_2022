package main
import (
	"fmt"
	"math"
)

func Y(x float64) float64{
	if (x >= 1){
		return math.Pow(1.2, x) - math.Pow(x, 1.2)
	}else{
		return math.Acos((x))
	}
}

func schet(start_x, end_x, delt_x float64){
	for i := start_x; i <= end_x; i += delt_x{
		fmt.Println((Y(i)))
	}
}

func schet_m(numbers [5]float64){
	for i := 0; i < len(numbers); i++{
		fmt.Println(Y(numbers[i]))
	}
}

func main(){
	fmt.Println("task A")
	schet(0.2, 2.2, 0.4)

	fmt.Println("\n task B")
	schet_m([5]float64{0.1, 0.9, 1.2, 1.5, 2.3})
}