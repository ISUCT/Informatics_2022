package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("ZADA")
	fmt.Println(internal.ZadA(1.25, 3.25, 0.4))
	fmt.Println("ZADB")
	fmt.Println(internal.ZadB([]float64{1.84, 2.71, 3.81, 4.56, 5.62}))
}
