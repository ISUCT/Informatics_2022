package main

import "fmt"

func main() {
	// type any array/list
	// Я не уверен, какой тип элементов должен быть в массиве, т.к не смог найти в го ни значения null ни underfined, про которые упоминается в условии
	array := [...]bool{true, false, true, true, true, false, false, true}
	fmt.Println(count_sheep(array[:]))
}

func count_sheep(sheep []bool) (count int) {
	count = 0
	for _, value := range sheep {
		if value == true {
			count += 1
		}
	}
	return
}
