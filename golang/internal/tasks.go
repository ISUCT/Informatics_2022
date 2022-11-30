package internal

import "math"

func formula(x float64, a float64, b float64) float64 {
	var y float64 = ((a*a*a)*(math.Cbrt(x)) - b*(math.Log10(x)/math.Log10(5))) / (math.Pow(math.Log10(x-1), 3))
	return y
}

func TaskA(a float64, b float64, xn float64, xk float64, xd float64) []float64 {
	
	var res []float64
	for i := xn; i < xk; i = i + xd {
		res = append(res, formula(i, a, b))
	}
	return res
}
func TaskB(a float64, b float64) []float64 {
	var x []float64 = []float64{1.9, 2.15, 2.34, 2.74, 3.16}
	var res []float64
	var i int8
	for ; i < 5; i++ {
		res = append(res, formula(x[i], a, b))
	}
	return res
}
