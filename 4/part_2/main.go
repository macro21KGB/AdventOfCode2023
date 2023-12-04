package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	input          string
	winningNumbers []int
	guessedNumbers []int
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

func (c Card) CalculateAmountOfCorrectGuesses() int {

	amountOfCorrectGuesses := 0
	for _, v := range c.guessedNumbers {
		for _, w := range c.winningNumbers {
			if v == w {
				amountOfCorrectGuesses++
			}
		}
	}
	return amountOfCorrectGuesses
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

func ElminateMultipleWhiteSpaces(input string) string {

	splitted := strings.Split(input, " ")
	results := make([]string, 0)
	for _, v := range splitted {
		if v != "" {
			results = append(results, v)
		}
	}

	return strings.Join(results, " ")
}

func UpdateMapWithDuplicatedCards(mapCard map[int][]Card, nextCardNumber, amountOfDuplicates int) map[int][]Card {

	maxMapCardNumber := len(mapCard)
	if nextCardNumber > maxMapCardNumber {
		return mapCard
	}

	// loop from 0 to amountOfDuplicates
	for i := 0; i < amountOfDuplicates; i++ {
		nextMapIndex := nextCardNumber + i
		if nextMapIndex > maxMapCardNumber {
			break
		}

		cards := mapCard[nextMapIndex]

		mapCard[cards[0].cardNumber] = append(mapCard[nextMapIndex], cards[0])
	}

	return mapCard
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
	mapCard := make(map[int][]Card)
	for scanner.Scan() {
		line := ElminateMultipleWhiteSpaces(scanner.Text())
		card := CreateNewCard(line)
		mapCard[card.cardNumber] = append(mapCard[card.cardNumber], card)
	}

	lengthOfMap := len(mapCard)

	for i := 1; i <= lengthOfMap; i++ {
		cards := mapCard[i]
		for _, card := range cards {
			amountOfCorrectGuesses := card.CalculateAmountOfCorrectGuesses()
			mapCard = UpdateMapWithDuplicatedCards(mapCard, card.cardNumber+1, amountOfCorrectGuesses)
		}

	}

	totalAmountOfCards := 0
	for _, cards := range mapCard {
		for range cards {
			totalAmountOfCards++
		}
	}

	log.Println("Total amount of cards: ", totalAmountOfCards)

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
