package main

func sheeps(slice []bool) int {
	nmofsheeps := 0
	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] == true {
			nmofsheeps++
		}
	}
	return nmofsheeps
}
