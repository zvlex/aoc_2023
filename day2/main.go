package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Game struct {
	id    int
	cubes []Cube
}

type Cube struct {
	green int
	blue  int
	red   int
}

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	games := []Game{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var g Game

		ns := strings.FieldsFunc(scanner.Text(), split)

		_, err := fmt.Sscanf(ns[0], "Game%d", &g.id)

		if err != nil {
			log.Fatal(err)
		}

		for _, n := range ns[1:] {
			c := Cube{}
			var value int
			var field string

			colors := strings.Split(n, ",")

			for _, color := range colors {
				fmt.Sscanf(color, "%d %s", &value, &field)

				switch field {
				case "green":
					c.green = value
				case "blue":
					c.blue = value
				case "red":
					c.red = value
				}
			}

			g.cubes = append(g.cubes, c)
		}

		games = append(games, g)
	}

	sum1 := partOne(games)
	sum2 := partTwo(games)

	fmt.Printf("Part One: %d\n", sum1)
	fmt.Printf("Part Two: %d\n", sum2)
}

func split(c rune) bool {
	return c == ':' || c == ';'
}

func partOne(games []Game) int {
	var sum int

	for _, g := range games {
		calc := true

		for _, c := range g.cubes {
			if c.red > maxRed || c.green > maxGreen || c.blue > maxBlue {
				calc = false
				break
			}
		}

		if calc {
			sum += g.id
		}
	}

	return sum
}

func partTwo(games []Game) int {
	var sum int

	for _, g := range games {

		var redPower int
		var greenPower int
		var bluePower int

		for _, c := range g.cubes {
			if redPower < c.red {
				redPower = c.red
			}

			if greenPower < c.green {
				greenPower = c.green
			}

			if bluePower < c.blue {
				bluePower = c.blue
			}

		}

		sum += redPower * greenPower * bluePower

	}

	return sum
}
