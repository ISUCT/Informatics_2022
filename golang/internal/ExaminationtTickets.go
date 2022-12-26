package internal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func AllTickets() {
	//Экзаменационные билеты
	Ticket1(9)
	fmt.Println(Ticket4([]float64{4, 1, 0.1, 100}))
	fmt.Println(Ticket5([]float64{4, 11, 10, 5, 8, 1}))
	fmt.Println(Ticket6(2))
	Ticket7(5, 2)
	fmt.Println(Ticket9(9))
	fmt.Println(Ticket10([3]int{1, 2, 3}, 10))
	Ticket12(2, 3)
	fmt.Println(Ticket14(10, "F"))
	fmt.Println(Ticket15(1000, "K"))
	fmt.Println(Ticket16(10))
	fmt.Println(Ticket17([]int8{0, 1, 1, 0, 1, 0}))
	Ticket18(9)
	fmt.Println(Ticket19(4, 4))
	fmt.Println(Ticket20(0, 1.5, 2, 10))
	fmt.Println(Ticket21([]int{1, 2, 3, 4}))
	fmt.Println(Ticket22([5]float64{1, 2, 3, 4, 5}))
	//CodeWars
	Sortedyesnohow([]int{1, 2, 3})
	Aretheysquare([]int{})
	fmt.Println(Fizzbuzz(5))
	fmt.Println(Likesvsdislikes([]string{"Like", "Like"}))
	fmt.Println(Moving("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(Insertdashes(41100130))
	Bingoornot([10]int{7, 2, 6, 1, 9, 14, 15, 8, 0, 4})
}

func Ticket1(num int) {
	fmt.Println("Экзаменационный билет №1")
	for i := 1; i < num+1; i++ {
		var array = []int{}
		for j := 1; j < 10; j++ {
			array = append(array, i*j)
		}
		fmt.Println(array)
	}
}

func Ticket4(array []float64) float64 {
	fmt.Println("Экзаменационный билет №4")
	var max_num float64 = array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max_num {
			max_num = array[i]
		}
	}
	return max_num
}

func Ticket5(array []float64) float64 {
	fmt.Println("Экзаменационный билет №5")
	var max_multi float64 = array[0] * array[1]
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]*array[j] > max_multi {
				max_multi = array[i] * array[j]
			}
		}
	}
	return max_multi
}

func Ticket6(b float64) []float64 {
	fmt.Println("Экзаменационный билет №6")
	var array = []float64{}
	for x := 0.0; x <= 3*math.Pi; x = x + 0.1*math.Pi {
		var y float64 = 1 + math.Sin(math.Pow(b, 3)+math.Pow(x, 3))
		array = append(array, y)
	}
	return array
}

func Ticket7(width int, height int) {
	fmt.Println("Экзаменационный билет №7")
	var width_string = "▉"
	for i := 1; i < width; i++ {
		width_string = width_string + "▉"
	}
	for i := 0; i < height; i++ {
		fmt.Println(width_string)
	}
}

func Ticket9(num int) []int {
	fmt.Println("Экзаменационный билет №9")
	var array = []int{}
	for i := 1; i < num+1; i++ {
		array = append(array, i)
	}
	for i := num - 2; i >= 1; i = i - 2 {
		array = append(array, i)
	}
	return array
}

func Ticket10(array [3]int, lim int) []int {
	fmt.Println("Экзаменационный билет №10")
	var num_array = []int{}
	for i := 0; i < 3; i++ {
		num_array = append(num_array, array[i])
	}
	for i := 0; i < lim; i++ {
		num_array = append(num_array, num_array[i]+num_array[i+1]+num_array[i+2])
	}

	return num_array
}

func Ticket12(player1 int, player2 int) {
	fmt.Println("Экзаменационный билет №12")
	if player1-player2 == -1 {
		fmt.Println("Победил первый игрок")
	} else if player1-player2 == -2 {
		fmt.Println("Победил второй игрок")
	} else if player1-player2 == 1 {
		fmt.Println("Победил второй игрок")
	} else if player1-player2 == 2 {
		fmt.Println("Победил первый игрок")
	} else {
		fmt.Println("Ничья")
	}
}

func Ticket14(temperature float64, convertTo string) float64 {
	fmt.Println("Экзаменационный билет №14")
	if strings.ToUpper(convertTo) == "C" {
		return (temperature - 32) / 1.8
	} else {
		return 1.8*temperature + 32
	}
}

// Сказано 2 метода, но сделаем 1
func Ticket15(distance float64, convertTo string) float64 {
	fmt.Println("Экзаменационный билет №15")
	if strings.ToUpper(convertTo) == "K" {
		return distance * 1.609
	} else {
		return distance / 1.609
	}
}

func Ticket16(num int) int {
	fmt.Println("Экзаменационный билет №16")
	var factorial int = 1
	for i := 1; i < num+1; i++ {
		factorial = factorial * i
	}
	return factorial
}
func Ticket17(array []int8) []int8 {
	fmt.Println("Экзаменационный билет №17")
	var coded_array = []int8{}
	if array[0] == 0 {
		coded_array = append(coded_array, 0)
	} else {
		coded_array = append(coded_array, 1)
	}
	for i := 1; i < len(array); i++ {
		if array[i] == array[i-1] {
			coded_array = append(coded_array, 0)
		} else {
			coded_array = append(coded_array, 1)
		}
	}
	return coded_array
}

func Ticket18(height int) {
	fmt.Println("Экзаменационный билет №18")
	var block string = "▉"
	for i := 0; i < height; i++ {
		fmt.Println(strings.Repeat(" ", height-i) + strings.Repeat(block, 1+2*i))
	}
}

func Ticket19(width int, height int) [][]int {
	fmt.Println("Экзаменационный билет №19")

	var result = [][]int{}
	for i := 1; i < (width*height + 1); i++ {
		var array1d = []int{}
		array1d = append(array1d, i)
		result = append(result, array1d)
	}
	return result
}

func Ticket20(x float64, a float64, b float64, n float64) []float64 {
	fmt.Println("Экзаменационный билет №20")
	var numbers = []float64{}
	for i := x; i <= 30; i = i + n {
		var y float64 = math.Cbrt(math.Pow(i, a) + b)
		numbers = append(numbers, y)
	}
	return numbers
}

func Ticket21(array []int) []int {
	fmt.Println("Экзаменационный билет №21")
	var result = []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 1 {
			result = append(result, array[i])
		}
	}
	return result
}

func Ticket22(array [5]float64) float64 {
	fmt.Println("Экзаменационный билет №22")
	var sum float64
	for i := 0; i < len(array); i++ {
		sum = sum + array[i]
	}
	var average float64 = sum / 5
	return average
}

func Sortedyesnohow(array []int) {
	fmt.Println("CodeWars 'Sorted? yes? no? how?'")
	var flag bool = true
	//Проверка на возрастание
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			flag = false
		}
	}
	if flag {
		fmt.Println("yes, ascending")
	} else {
		flag = true
		for i := 1; i < len(array); i++ {
			if array[i] > array[i-1] {
				flag = false
			}
		}
		if flag {
			fmt.Println("yes, descending")
		} else {
			fmt.Println("no")
		}
	}
}

func Aretheysquare(array []int) {
	fmt.Println("CodeWars 'Are they square?'")
	var massivflag int = 0
	if len(array) == 0 {
		massivflag = 2
	} else {
		for i := 0; i < len(array); i++ {
			var numflag bool = false
			for j := 0; j < array[i]; j++ {
				if j*j == array[i] {
					numflag = true
				}
			}
			if numflag {
				massivflag = 1
			}
		}
	}
	if massivflag == 1 {
		fmt.Println(true)
	} else if massivflag == 0 {
		fmt.Println(false)
	} else {
		fmt.Println(nil)
	}
}

func Fizzbuzz(n int) []string {
	fmt.Println("CodeWars 'Fizz Buzz'")
	var numbers_array = []int{}
	var array = []string{}
	for i := 1; i < n+1; i++ {
		numbers_array = append(numbers_array, i)
	}
	for i := 0; i < len(numbers_array); i++ {
		if numbers_array[i]%3 == 0 {
			array = append(array, "Fizz")
		} else if numbers_array[i]%5 == 0 {
			array = append(array, "Buzz")
		} else {
			array = append(array, fmt.Sprint(numbers_array[i]))
		}
	}
	return array
}

func Likesvsdislikes(array []string) string {
	fmt.Println("CodeWars 'Likes Vs Dislikes'")
	var Likes bool = false
	var Dislikes bool = false
	for i := 0; i < len(array); i++ {
		switch array[i] {
		case "Dislike":
			if Dislikes {
				Dislikes = false
			} else {
				Dislikes = true
			}
			if Likes {
				Likes = false
			}
		case "Like":
			if Likes {
				Likes = false
			} else {
				Likes = true
			}
			if Dislikes {
				Dislikes = false
			}

		default:
		}
	}

	if Likes {
		return "Like"
	}
	if Dislikes {
		return "Dislike"
	} else {
		return "Nothing"
	}
}

func Moving(s string) string {
	fmt.Println("CodeWars 'Move 10'")
	var array = []byte{}
	var fstr string
	for i := 0; i < len(s); i++ {
		array = append(array, s[i])
	}
	for i := 0; i < len(array); i++ {
		if int(array[i])+10 > 122 {
			fstr = fstr + string(array[i]+10-122+96)
		} else {
			fstr = fstr + string(array[i]+10)
		}
	}
	return fstr
}

func Insertdashes(num int) string {
	fmt.Println("CodeWars 'Insert dashes'")
	var fstr string
	var numstr = strings.Split(fmt.Sprint(num), "")
	var numbers = []int{}
	for i := 0; i < len(numstr); i++ {
		var res, err = strconv.ParseInt(numstr[i], 10, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, int(res))
	}
	var fnumbers = []string{fmt.Sprint(numbers[0])}
	for i := 1; i < len(numbers); i++ {
		if numbers[i]%2 == 1 && numbers[i-1]%2 == 1 {
			fnumbers = append(fnumbers, "-")
			fnumbers = append(fnumbers, fmt.Sprint(numbers[i]))
		} else {
			fnumbers = append(fnumbers, fmt.Sprint(numbers[i]))
		}
	}
	for i := 0; i < len(fnumbers); i++ {
		fstr = fstr + fnumbers[i]
	}
	return fstr
}

func Bingoornot(array [10]int) {
	fmt.Println("CodeWars 'Bingo ( Or Not )'")
	var counter = 0
	var flag2 bool = false
	var flag7 bool = false
	var flag9 bool = false
	var flag14 bool = false
	var flag15 bool = false
	for i := 0; i < len(array); i++ {
		switch array[i] {
		case 2:
			flag2 = true
		case 7:
			flag7 = true
		case 9:
			flag9 = true
		case 14:
			flag14 = true
		case 15:
			flag15 = true
		default:
		}
	}
	if flag2 {
		counter++
	}
	if flag7 {
		counter++
	}
	if flag9 {
		counter++
	}
	if flag14 {
		counter++
	}
	if flag15 {
		counter++
	}
	if counter == 5 {
		fmt.Println("WIN")
	} else {
		fmt.Println("LOSE")
	}
}
