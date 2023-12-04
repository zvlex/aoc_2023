package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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


	sum1 := partOne(lines)
	fmt.Printf("Part One: %d\n", sum1)

	sum2 := partTwo(lines)
	fmt.Printf("Part Two: %d\n", sum2)
}

func partOne(lines []string) int {
	return calculateSum(lines)
}

func partTwo(lines []string) int {
	var newLines []string

	replacers := []*strings.Replacer {
			strings.NewReplacer("one", "o1e"),
			strings.NewReplacer("two", "t2o"),
			strings.NewReplacer("three", "t3e"),
			strings.NewReplacer("four", "4"),
			strings.NewReplacer("five", "f5e"),
			strings.NewReplacer("six", "6"),
			strings.NewReplacer("seven", "s7n"),
			strings.NewReplacer("eight", "e8t"),
			strings.NewReplacer("nine", "n9e"),
	}

	for _, l := range lines {
		nl := applyReplacers(l, replacers...)
		newLines = append(newLines, nl)
	}

	return calculateSum(newLines)
}

func applyReplacers(line string, replacers ...*strings.Replacer) string {
	result := line
	for _, r := range replacers {
		result = r.Replace(result)
	}
	return result
}

func calculateSum(lines []string) int {
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

	return sum
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
