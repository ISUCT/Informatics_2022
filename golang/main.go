package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func formula(x float64) float64 {
	var a float64 = -2.5
	var b float64 = 3.4
	var i float64 = 0
	if x > 5 {
		var y float64 = (3 * math.Log10(a*a+x)) / ((a + x) * (a + x))
		i = y
	}
	if x < 6 {
		var y float64 = math.Pow((a+b*x), 2.5) / (1.8 + math.Pow(math.Cos(a*x), 3))
		i = y
	}
	return i
}

func mainn() {
	fmt.Println("Задание А")
	var x float64 = 3.5
	for ; x < 6.5; x = x + 0.6 {
		fmt.Println(formula(x))
	}
	fmt.Println("Задание В")
	fmt.Println(2.89)
	fmt.Println(3.54)
	fmt.Println(5.21)
	fmt.Println(6.28)
	fmt.Println(3.47)
}

func task1(n int64) {
	if n%2 == 0 {
		fmt.Println("Even")
	}
	if n%2 != 0 {
		fmt.Println("Odd")
	}
}

func task2() {
	sheeps := [...]bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	kol := 0
	for i := 0; i < len(sheeps); i++ {
		if sheeps[i] == true || sheeps[i] == false {
			if sheeps[i] == true {
				kol += 1
			}
		} else {
			fmt.Println("Values ", i, " undefined")
		}
	}
	fmt.Println(kol)
}

func task3() {
	n := 0
	fmt.Scan(&n)
	monkeys := make([]int, n)
	for i := 0; i < n; i++ {
		monkeys[i] = i + 1
	}
	fmt.Println(monkeys)
}

func task4() {
	n := 0
	m := 0
	fmt.Println("Classmates = ")
	fmt.Scan(&n)
	fmt.Println("Pages = ")
	fmt.Scan(&m)
	if m > 0 && n > 0 {
		kol := m * n
		fmt.Println(kol)
	} else {
		fmt.Println("0")
	}

}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func task5() {
	bullets := 0
	fmt.Println("Bullets: ")
	fmt.Scan(&bullets)
	rand.Seed(time.Now().UnixNano())
	dragons := randomInt(1, 11)
	if bullets >= (dragons * 2) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	fmt.Println("There were ", dragons, " dragons")
}

func main() {
	fmt.Println("Задание №1")
	x := 0
	fmt.Print("x = ")
	fmt.Scan(&x)
	task1(int64(x))
	fmt.Println("Задание №2")
	task2()
	fmt.Println("Залание №3")
	fmt.Println("Введите количесвто обезьян")
	task3()
	fmt.Println("Задание №4")
	task4()
	fmt.Println("Задание №5")
	task5()

}
