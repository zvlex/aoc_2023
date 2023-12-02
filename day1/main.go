package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var sum int

	for _, line := range lines {

		left := findLeft(line)
		right := findRight(line)

		currentDigit := (left + right)

		d, err := strconv.Atoi(currentDigit)
		if err != nil {
			log.Fatal(err)
		}

		sum += d
	}

	fmt.Println(sum)
}

func startIndex(index int) func(int, int) int {
	return func(lineLen, el int) int {
		if index < 0 {
			el = (lineLen - el + index)
		}

		return el
	}
}

func findRight(line string) string {
	digit := findDigitSymbol(line, startIndex(-1))

	return string(digit)
}

func findLeft(line string) string {
	digit := findDigitSymbol(line, startIndex(0))

	return string(digit)
}

func findDigitSymbol(line string, calcIndex func(int, int) int) rune {
	var symbol rune

	lineLen := len(line)

	for i := range line {
		ch_el := rune(line[calcIndex(lineLen, i)])

		if isDigit(ch_el) {
			symbol = ch_el
			break
		}
	}

	return symbol
}

func isDigit(ch rune) bool {
	return ch > 46 && ch < 58
}
