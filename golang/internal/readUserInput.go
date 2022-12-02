package internal

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadUserInput() []float64 {
	resultSlice := make([]float64, 0, 20)
	// Да, есть разница между сканнером и ридером, но я решила выбрать сканнер, так как
	// его относительно проще использовать, хоть и существует ограничение на количество
	// строк (64 тысячи); плюс сканнер не возвращает символ новой строки в конце
	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		sliceScan := reader.Text()                        // Получение текста со сканнера
		splitValuesSlice := strings.Split(sliceScan, " ") // Разбиение введёного текста на части (на выходе получается слайс)

		for _, element := range splitValuesSlice {
			fl64, err := strconv.ParseFloat(element, 64) // Конвертация строки в float64
			if err != nil {
				log.Fatal("Incorrect value for strconv.ParseFloat (readUserInput.go)")
			}
			resultSlice = append(resultSlice, fl64)
		}
	}
	if reader.Err() != nil {
		log.Fatal("Incorrect input of x values")
	}

	return resultSlice
}
