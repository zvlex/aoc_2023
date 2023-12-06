package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Card struct {
	id int
	cardNums []string
	yNums []string
}

type WinCard struct {
	points int
	winNums []string
}


func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cards []Card

	for scanner.Scan() {
		var card Card

		data := strings.FieldsFunc(scanner.Text(), split)

		_, err := fmt.Sscanf(data[0], "Card%d", &card.id)

		if err != nil {
			log.Fatal(err)
		}


		card.cardNums = append(card.cardNums, strings.Fields(data[1])...)
		card.yNums = append(card.yNums, strings.Fields(data[2])...)

		cards = append(cards, card)
	}

	var winCards []WinCard

	for _, card := range cards {
		var wCard WinCard

		factor := 1

		for _, cN := range card.cardNums {
			var found bool

			yLen := len(card.yNums)

			for j, yN := range card.yNums {
				//fmt.Println(factor)
				if cN == yN {
					if wCard.points == 0 {
						wCard.points = 1
					}
					wCard.points *= factor
					wCard.winNums = append(wCard.winNums, cN)

					factor = 2
					found = true

					//
					//fmt.Println("----->")
					//fmt.Println(">", yN, cN, wCard.points)
					break
				}

				if found && yLen - 1 == j {
					factor = 1
				}


			}
		}


		if wCard.points > 0 {
			winCards = append(winCards, wCard)
		}
	}


	sum1 := partOne(winCards)
	fmt.Println("Part One: ", sum1)
}

func partOne(winCards []WinCard) int {
	var sum int

	for _, c := range winCards {
		sum += c.points
	}

	return sum
}


func split(c rune) bool {
	return c == ':' || c == '|'
}
