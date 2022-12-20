package internal

import "math"

func ZadA(xn float64, xk float64, dx float64) []float64 {
	var x float64 = xn
	var massivA = make([]float64, 0)
	for ; x <= xk; x += dx {
		massivA = append(massivA, form(x))
	}
	return massivA
}

func ZadB(massivB []float64) []float64 {
	var i int8 = 0
	for ; i <= 4; i += 1 {
		massivB[i] = form(massivB[i])
	}
	return massivB
}

func form(y float64) float64 {
	return (math.Sqrt(math.Sqrt(math.Abs(y*y-2.5))) + math.Cbrt(math.Log10(y*y)))
}
