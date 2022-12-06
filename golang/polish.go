package main

import (
	"strings"
)

func polish(inps string) string {
	pa := []string{"ą", "ć", "ę", "ł", "ń", "ó", "ś", "ź", "ż"}
	ea := []string{"a", "c", "e", "l", "n", "o", "s", "z", "z"}
	inpscopy := inps
	for i := 0; i < len(inps); i++ {
		for j := 0; j < len(pa); j++ {
			inpscopy = strings.ReplaceAll(inpscopy, pa[j], ea[j])
		}
	}
	return inpscopy
}
