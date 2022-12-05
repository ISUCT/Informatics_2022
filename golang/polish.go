package main

import (
	"strings"
)

func polish(inps string) string {
	pa := []string{"ą", "ć", "ę", "ł", "ń", "ó", "ś", "ź", "ż"}
	ea := []string{"a", "c", "e", "l", "n", "o", "s", "z", "z"}
	psl := strings.Split(inps, "")[:]
	for i := 0; i <= len(inps)-1; i++ {
		for j := 0; j <= len(pa)-1; j++ {
			if psl[i] == pa[j] {
				psl[i] = ea[j]
			}
		}
	}
	engstr := ""
	for k := 0; k <= len(psl)-1; k++ {
		engstr = engstr + psl[k]
	}
	return engstr
}
