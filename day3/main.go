package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Part struct {
	row      int
	col      int
	symbol   string
	values   []string
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var table = [][]string{}

	for scanner.Scan() {
		var row []string

		for _, t := range scanner.Text() {
			row = append(row, string(t))
		}

		table = append(table, row)
	}

	var parts []Part

	for i := 0; i < len(table); i++ {
		row := table[i]

		var sNum []string
		var t bool

		var p Part

		for j := 0; j < len(row); j++ {
			if row[j] == "." || isSymbol(row[j]) {
				if t {
					p.values = sNum
					parts = append(parts, p)
					t = false
				}
				sNum = nil
				continue
			}

			top := i - 1
			bottom := i + 1

			prev := j - 1
			next := j + 1

			if top < 0 {
				top = 0
			}

			if bottom > len(table)-1 {
				bottom = len(table) - 1
			}

			if prev < 0 {
				prev = 0
			}

			if next > len(row)-1 {
				next = len(row) - 1
			}

			if isDigit(row[j]) {
				if isSymbol(row[prev]) {
					p.row = prev
					p.col = i
					p.symbol = row[prev]
					t = true
				}

				if isSymbol(row[next]) {
					p.row = next
					p.col = i
					p.symbol = row[next]

					t = true
				}

				indices := []struct{ i, j int }{
					{top, prev}, {top, next},
					{bottom, prev}, {bottom, next},
					{top, j}, {bottom, j},
				}

				for _, index := range indices {
					result := table[index.i][index.j]

					if isSymbol(result) {
						p.row = index.j
						p.col = index.i
						p.symbol = result
						t = true

						break
					}
				}

				sNum = append(sNum, row[j])


				if j == len(row)-1 && t {
					p.values = sNum
					parts = append(parts, p)
					t = false

					sNum = nil
				}

			}
		}
	}

	sum1 := partOne(parts)
	fmt.Printf("Part One: %d\n", sum1)

	sum2 := partTwo(parts)
	fmt.Printf("Part Two: %d\n", sum2)
}

func partOne(parts []Part) int {
	var sum int

	for _, part := range parts {
		result := strings.Join(part.values, "")

		d, err := strconv.Atoi(result)

		if err != nil {
			log.Fatal(err)
		}

		sum += d
	}

	return sum
}

func partTwo(parts []Part) int {
	var sum int

	k := make(map[string]int)

	for _, p1 := range parts {
		for _, p2 := range parts {
			equal := reflect.DeepEqual(p1.values, p2.values)

			e := p1.row == p2.row && p1.col == p2.col
			if  e && equal  {
				continue
			}

			if e && p1.symbol == p2.symbol && p1.symbol == "*" && p2.symbol == "*" {
				result1 := strings.Join(p1.values, "")

				d1, err := strconv.Atoi(result1)

				if err != nil {
					log.Fatal(err)
				}

				result2 := strings.Join(p2.values, "")

				d2, err := strconv.Atoi(result2)

				if err != nil {
					log.Fatal(err)
				}

				key := strconv.Itoa(p1.row) + p1.symbol + strconv.Itoa(p1.col)

				k[key] = d1 * d2


			}
		}
	}


	for _, v := range k {
		sum += v
	}

	return sum
}

func isDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func isSymbol(s string) bool {
	if s == "." {
		return false
	}

	res, _ := regexp.MatchString(`\D`, s)
	return res
}
