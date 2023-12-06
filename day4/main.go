package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Card struct {
	id       int
	cardNums []string
	yNums    []string
}

type WinCard struct {
	cardId    int
	points    int
	quantity  int
	winNums   []string
	matchNums int
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
				if cN == yN {
					if wCard.points == 0 {
						wCard.points = 1
					}
					wCard.points *= factor
					wCard.winNums = append(wCard.winNums, cN)
					wCard.matchNums += 1

					factor = 2
					found = true

					break
				}

				if found && yLen-1 == j {
					factor = 1
				}

			}
		}

		wCard.quantity = 1
		wCard.cardId = card.id
		winCards = append(winCards, wCard)
	}

	sum1 := partOne(winCards)
	fmt.Println("Part One: ", sum1)

	sum2 := partTwo(winCards)
	fmt.Println("Part Two: ", sum2)
}

func partOne(winCards []WinCard) int {
	var sum int

	for _, c := range winCards {
		sum += c.points
	}

	return sum
}

func partTwo(winCards []WinCard) int {
	var sum int

	for i, wc := range winCards {
		for j := 0; j < wc.matchNums; j++ {
			winCards[i+j+1].quantity += wc.quantity
		}
	}

	for _, wc := range winCards {
		sum += wc.quantity
	}

	return sum
}

func split(c rune) bool {
	return c == ':' || c == '|'
}
