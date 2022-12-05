package main

func hero(bullets, dragons int) bool {
	if bullets/dragons < 2 {
		return false
	}
	return true
}
