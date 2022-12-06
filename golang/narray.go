package main

func narray(inpsl []int, n int) []int { //Я не знаю, как передать массив с неизвестным изначально количеством аргументов. Поэтому использую slice
	newsl := []int{}
	for i := 0; i < len(inpsl); i++ {
		if inpsl[i] == n {
			newsl = append(newsl, i)
		}
	}
	if len(newsl) > 0 {
		return newsl
	} else {
		return []int{0}
	}

}
