package main

func monkeys(n int) []int {
	nofm := []int{}
	for i := 1; i <= n; i++ {
		nofm = append(nofm, i)
	}
	return nofm
}
