package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	input          string
	winningNumbers []int
	guessedNumbers []int
	score          int
	cardNumber     int
}

func ConvertStringArrayToInt(arr []string) []int {

	results := make([]int, 0)
	for _, v := range arr {
		convertedValue, error := strconv.Atoi(v)
		if error == nil {
			results = append(results, convertedValue)
		}

	}
	return results
}

func (c *Card) ParseInput() {

	splitted := strings.Split(c.input, ":")
	c.cardNumber, _ = strconv.Atoi(strings.Split(splitted[0], " ")[1])

	guessedAndWinningNumbers := strings.Split(splitted[1], "|")

	parsedWin := ConvertStringArrayToInt(strings.Split(guessedAndWinningNumbers[0], " "))
	parsedGuess := ConvertStringArrayToInt(strings.Split(guessedAndWinningNumbers[1], " "))

	c.winningNumbers = parsedWin
	c.guessedNumbers = parsedGuess

}

func CreateNewCard(input string) Card {

	newCard := Card{
		input:          input,
		winningNumbers: make([]int, 0),
		guessedNumbers: make([]int, 0),
		cardNumber:     0,
	}

	newCard.ParseInput()

	return newCard
}

func main() {

	// First Star
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	cards := make([]Card, 0)
	for scanner.Scan() {
		line := scanner.Text()
		card := CreateNewCard(line)
		cards = append(cards, card)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
