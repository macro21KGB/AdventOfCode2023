package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	Cards    map[string]int
	RawInput string
	Bid      int
}

type HandType int

// 0 is the highest power, 6 is the lowest
const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func CalculatePower(cardValues ...string) []int {

	ScaleMap := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}

	result := make([]int, 0)

	for _, value := range cardValues {
		result = append(result, ScaleMap[value])
	}

	return result
}

/*
Give a map of cards, group them by amount of them, for example
map[2:1 3:2 K:1 T:1] means that are 2 cards with value 3, 1 card with value 2, 1 card with value K and 1 card with value T
this function must group it by amount, so the result will be
map[1:2 2:1 3:1]
*/
func CheckHandType(cards map[string]int) HandType {

	// Check if there is a five of a kind
	for _, value := range cards {
		if value == 5 {
			return FiveOfAKind
		}

	}

	// Check if there is a four of a kind
	for _, value := range cards {
		if value == 4 {
			return FourOfAKind
		}

	}

	// Check if there is a three of a kind
	for _, value := range cards {
		if value == 3 {
			return ThreeOfAKind
		}

	}

	// Check if there is a two pair
	pairs := 0
	for _, value := range cards {
		if value == 2 {
			pairs++
		}
	}

	if pairs == 2 {
		return TwoPair
	}

	if pairs == 1 {
		return OnePair
	}

	return HighCard
}

func CreateHandDeck(input []string) map[string]int {
	cards := make(map[string]int, 0)

	for _, item := range input {
		_, ok := cards[item]
		if ok {
			cards[item] = cards[item] + 1
		} else {
			cards[item] = 1
		}
	}

	return cards
}

func SumHandPower(rawHandValues string) int {
	powers := CalculatePower(strings.Split(rawHandValues, "")...)

	result := 0
	for _, value := range powers {
		result += value
	}

	return result
}

func IsFirstHandMorePowerful(firstHand string, secondHand string) bool {

	firstHandPower := CalculatePower(strings.Split(firstHand, "")...)
	secondHandPower := CalculatePower(strings.Split(secondHand, "")...)

	for index, value := range firstHandPower {
		if value > secondHandPower[index] {
			return true
		}
	}

	return false
}

func main() {
	file, _ := os.Open("../input.txt")

	scanner := bufio.NewScanner(file)

	hands := make([]Hand, 0)
	for scanner.Scan() {
		line := scanner.Text()

		splittedLine := strings.Split(line, " ")
		bidNumber, _ := strconv.Atoi(splittedLine[1])

		hands = append(hands, Hand{
			Cards:    CreateHandDeck(strings.Split(strings.TrimSpace(splittedLine[0]), "")),
			RawInput: strings.TrimSpace(splittedLine[0]),
			Bid:      bidNumber,
		})

	}

	error := scanner.Err()
	if error != nil {
		log.Fatal(error)
	}

	// Part 1
	currentVictoryRank := 1
	mappedHands := make(map[HandType][]Hand, 0)
	for _, hand := range hands {
		mappedHands[CheckHandType(hand.Cards)] = append(mappedHands[CheckHandType(hand.Cards)], hand)
	}

	finalResult := 0
	for rank := 1; rank <= len(hands); rank++ {

		handsInMap, ok := mappedHands[HandType(rank)]
		if !ok {
			continue
		}

		// if there is only one hand, assign the rank directly
		if len(handsInMap) == 1 {
			finalResult += handsInMap[0].Bid * currentVictoryRank
			currentVictoryRank++
		}

		// if there is a conflict, check for the most powerful hand
		if len(handsInMap) > 1 {
			mostLowPowerHand := handsInMap[0]
			mostLowPowerHandIndex := 0

			for i := 0; i < len(handsInMap)-1; i++ {
				for j := i + 1; j < len(handsInMap); j++ {
					if IsFirstHandMorePowerful(handsInMap[i].RawInput, handsInMap[j].RawInput) {
						mostLowPowerHand = handsInMap[j]
						mostLowPowerHandIndex = j
					}
				}

				// remove the most low power hand from the list
				handsInMap = append(handsInMap[:mostLowPowerHandIndex], handsInMap[mostLowPowerHandIndex+1:]...)
				// add the bid to the final result
				finalResult += mostLowPowerHand.Bid * currentVictoryRank
				currentVictoryRank++
			}

			if len(handsInMap) == 1 {
				finalResult += handsInMap[0].Bid * currentVictoryRank
				currentVictoryRank++
			}
		}
	}

	fmt.Println(finalResult)
}
