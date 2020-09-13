package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	rows := readPrices(os.Args[1])
	start, end, col, percent := userInp()
	if start < 1 || end > (len(rows)-1) {
		println("Incorect input. Program closed")
		return
	}
	rows = addPersent(rows, start, end, col, percent)
	writePrices(os.Args[1], rows)
}

func readPrices(fname string) [][]string {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", fname, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}

func userInp() (int, int, int, float64) {
	fmt.Print("Введите № колонки где нужно поменять цены: ")
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatalln("Cannot read input string:", err.Error())
	}
	col, err := strconv.Atoi(strings.Replace(input, "\n", "", -1))
	col -= 1
	if err != nil {
		log.Fatalln("Cannot convert string to number:", err.Error())
	}

	fmt.Print("Введите № ряда с которого начать изьенения: ")
	input, err = r.ReadString('\n')
	if err != nil {
		log.Fatalln("Cannot read input string:", err.Error())
	}
	start, err := strconv.Atoi(strings.Replace(input, "\n", "", -1))
	start -= 1
	if err != nil {
		log.Fatalln("Cannot convert string to number:", err.Error())
	}

	fmt.Print("Введите № ряда на котором закончить изменения: ")
	input, err = r.ReadString('\n')
	if err != nil {
		log.Fatalln("Cannot read input string:", err.Error())
	}
	end, err := strconv.Atoi(strings.Replace(input, "\n", "", -1))
	end -= 1
	if err != nil {
		log.Fatalln("Cannot convert string to number:", err.Error())
	}

	fmt.Print("Введите процент на который изменить цену: ")
	input, err = r.ReadString('\n')
	if err != nil {
		log.Fatalln("Cannot read input string:", err.Error())
	}
	percent, err := strconv.ParseFloat(strings.Replace(input, "\n", "", -1), 64)
	if err != nil {
		log.Fatalln("Cannot convert string to float32:", err.Error())
	}

	return start, end, col, percent
}

func addPersent(rows [][]string, start, end, col int, percent float64) [][]string {
	for i := start; i <= end; i++ {
		fmt.Printf("price: %v i: %d\n", rows[i][col], i)
		price, err := strconv.ParseFloat(rows[i][col], 64)
		if err != nil {
			continue
			log.Fatalln("Cannot convert string to float64:", err.Error())
		}

		updatedPrice := int64(math.Round(price + ((price / 100) * percent)))
		rows[i][col] = strconv.FormatInt(updatedPrice, 10)
		fmt.Printf("updated price: %v %d\n", rows[i][col], i)
	}

	return rows
}

func writePrices(fileName string, rows [][]string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", fileName, err.Error())
	}

	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close '%s': %s\n", fileName, e.Error())
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}
