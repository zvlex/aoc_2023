package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var labels = map[string]string{
	"A": "60",
	"K": "50",
	"Q": "40",
	"J": "30",
	"T": "20",
	"9": "19",
	"8": "18",
	"7": "17",
	"6": "16",
	"5": "15",
	"4": "14",
	"3": "13",
	"2": "12",
}

type Card struct {
	hand  string
	bid   int
	props HandProps
}

type HandProps struct {
	weight int
	score  int
	label  string
}

func fetchHandData(hands string) HandProps {
	res := map[string]int{}
	var weights []string

	for _, v := range hands {
		s := string(v)

		res[s] += 1

		value, exists := labels[s]

		if !exists {
			value = s
		}

		weights = append(weights, value)
	}

	sameLabelsQuantity := 1
	uniqLabelsQuantity := len(res)

	for _, v := range res {
		if v > sameLabelsQuantity {
			sameLabelsQuantity = v
		}
	}

	var handProps HandProps

	weight := strings.Join(weights, "")

	w, err := strconv.Atoi(weight)
	if err != nil {
		log.Fatal(err)
	}

	handProps.weight = w

	switch {
	case sameLabelsQuantity == 5:
		handProps.score = 10
		handProps.label = "Five of a kind"
	case sameLabelsQuantity == 4 && uniqLabelsQuantity == 2:
		handProps.score = 9
		handProps.label = "Four of a kind"
	case sameLabelsQuantity == 3 && uniqLabelsQuantity == 2:
		handProps.score = 8
		handProps.label = "Full house"
	case sameLabelsQuantity == 3 && uniqLabelsQuantity == 3:
		handProps.score = 7
		handProps.label = "Three of a kind"
	case sameLabelsQuantity == 2 && uniqLabelsQuantity == 3:
		handProps.score = 6
		handProps.label = "Two pair"
	case sameLabelsQuantity == 2 && uniqLabelsQuantity == 4:
		handProps.score = 5
		handProps.label = "One pair"
	default:
		handProps.score = 1
		handProps.label = "High card"
	}

	return handProps
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
		res := strings.Fields(scanner.Text())

		handProps := fetchHandData(res[0])

		bid, err := strconv.Atoi(res[1])
		if err != nil {
			log.Fatal(err)
		}

		card := Card{hand: res[0], bid: bid, props: handProps}

		cards = append(cards, card)
	}

	sort.Slice(cards, func(x, y int) bool {
		propsX := cards[x].props
		propsY := cards[y].props

		if propsX.score != propsY.score {
			return propsX.score > propsY.score
		}

		return propsX.weight > propsY.weight
	})

	sum1 := partOne(cards)
	fmt.Println("Part One: ", sum1)
}

func partOne(cards []Card) int {
	var sum int

	rank := len(cards)

	for _, cX := range cards {
		sum += cX.bid * rank
		rank -= 1
	}

	return sum
}
