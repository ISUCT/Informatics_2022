package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func main() {
	fmt.Println("Even or odd:", internal.Even_or_odd(10))
	fmt.Println("Counting sheep:", internal.Sheep([]bool{
		true, true, true, false, false, true, true,
	}))
	fmt.Println("Monkey count:", internal.count_with_son(15))
	fmt.Println("School:", internal.School(5, 5))
	fmt.Println("Hero:", internal.Lucky_hero(9, 4))
	fmt.Println("Polish messages:", internal.Pol_to_eng("Jędrzej Błądziński"))
	fmt.Println("7:", internal.Find_all([]int(6, 9, 3, 4, 3, 82, 4), 3))
	fmt.Println("8:", internal.Only_min([][]int{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 35, 56, 100}))
}
