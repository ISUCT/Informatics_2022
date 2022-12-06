package main

import "fmt"

func main() {
	var monkeys_k int
	fmt.Print("How much monkeys there? > ")
	fmt.Scanln(&monkeys_k)
	fmt.Println(generate_array(monkeys_k))
}

func generate_array(k int) (counting []int) {
	for i := 1; i <= k; i++ {
		counting = append(counting, i)
	}
	return
}
