package main

import (
	"bufio"
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

func CalculatePower(cardValue string) int {

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

	return ScaleMap[cardValue]

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

func main() {
	file, _ := os.Open("../test_input.txt")

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
	log.Println(hands)

}
