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

var labels = map[string]int{
	"2": 12,
	"3": 13,
	"4": 14,
	"5": 15,
	"6": 16,
	"7": 17,
	"8": 18,
	"9": 19,
	"J": 30,
	"T": 20,
	"Q": 40,
	"K": 50,
	"A": 60,
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

var joker = "J"
var jack = "J"

var replaceRules = fetchReplaceRules()

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cards1 []Card
	var cards2 []Card

	for scanner.Scan() {
		text := scanner.Text()

		cards1 = append(cards1, parseCard(text, fetchHandData))
		cards2 = append(cards2, parseCard(text, fetchHandData2))
	}

	cards1 = sortCards(cards1)
	cards2 = sortCards(cards2)

	sum1 := calculateSum(cards1)
	fmt.Println("Part One: ", sum1)

	sum2 := calculateSum(cards2)
	fmt.Println("Part Two: ", sum2)
}

func fetchHandData(hand string) HandProps {
	res := map[string]int{}

	sameLabelQuantity := 1

	for k := range labels {
		labelQuantity := strings.Count(hand, k)

		if labelQuantity > 0 {
			res[k] = labelQuantity

			if labelQuantity > sameLabelQuantity {
				sameLabelQuantity = labelQuantity
			}
		}
	}

	currentReplaceRules := append(replaceRules, jack, "30")

	uniqLabelQuantity := len(res)

	replacer := strings.NewReplacer(currentReplaceRules...)

	weight := replacer.Replace(hand)

	return generateHandProps(weight, sameLabelQuantity, uniqLabelQuantity)
}

func fetchHandData2(hand string) HandProps {
	res := map[string]int{}

	sameLabelQuantity := 1

	for k := range labels {
		labelQuantity := strings.Count(hand, k)

		if labelQuantity > 0 {
			res[k] = labelQuantity

			if labelQuantity > sameLabelQuantity && k != joker {
				sameLabelQuantity = labelQuantity
			}
		}
	}

	if res[joker] < 5 {
		sameLabelQuantity += res[joker]
	} else {
		sameLabelQuantity = res[joker]
	}

	currentReplaceRules := append(replaceRules, joker, "10")

	delete(res, joker)

	uniqLabelQuantity := len(res)

	replacer := strings.NewReplacer(currentReplaceRules...)

	weight := replacer.Replace(hand)

	return generateHandProps(weight, sameLabelQuantity, uniqLabelQuantity)
}

func generateHandProps(wStr string, sameLabelQuantity int, uniqLabelQuantity int) HandProps {
	var handProps HandProps

	w, err := strconv.Atoi(wStr)
	if err != nil {
		log.Fatal(err)
	}

	handProps.weight = w

	switch {
	case sameLabelQuantity == 5:
		handProps.score = 10
		handProps.label = "Five of a kind"
	case sameLabelQuantity == 4 && uniqLabelQuantity == 2:
		handProps.score = 9
		handProps.label = "Four of a kind"
	case sameLabelQuantity == 3 && uniqLabelQuantity == 2:
		handProps.score = 8
		handProps.label = "Full house"
	case sameLabelQuantity == 3 && uniqLabelQuantity == 3:
		handProps.score = 7
		handProps.label = "Three of a kind"
	case sameLabelQuantity == 2 && uniqLabelQuantity == 3:
		handProps.score = 6
		handProps.label = "Two pair"
	case sameLabelQuantity == 2 && uniqLabelQuantity == 4:
		handProps.score = 5
		handProps.label = "One pair"
	default:
		handProps.score = 1
		handProps.label = "High card"
	}

	return handProps
}

func fetchReplaceRules() []string {
	var rules []string

	for k, v := range labels {
		if k == joker || k == jack {
			continue
		}

		rules = append(rules, k)
		rules = append(rules, fmt.Sprint(v))
	}

	return rules
}

func sortCards(cards []Card) []Card {
	sort.Slice(cards, func(x, y int) bool {
		propsX := cards[x].props
		propsY := cards[y].props

		if propsX.score != propsY.score {
			return propsX.score > propsY.score
		}

		return propsX.weight > propsY.weight
	})

	return cards
}

func parseCard(value string, fetchData func(string) HandProps) Card {
	res := strings.Fields(value)

	handProps := fetchData(res[0])

	bid, err := strconv.Atoi(res[1])
	if err != nil {
		log.Fatal(err)
	}

	return Card{hand: res[0], bid: bid, props: handProps}
}

func calculateSum(cards []Card) int {
	var sum int

	rank := len(cards)

	for _, c := range cards {
		//fmt.Println(c.hand, c.props.label, rank, c.props.score, c.props.weight)
		sum += c.bid * rank
		rank -= 1
	}

	return sum
}
