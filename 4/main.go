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
}

func ConvertStringArrayToInt(arr []string) ([]int, error) {
	result := make([]int, len(arr))
	for i, s := range arr {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result[i] = num
	}
	return result, nil
}

func (c *Card) ParseInput() {
	splitted := strings.Split(c.input, ":")
	guessedAndWinningNumbers := strings.Split(splitted[1], "|")

	parsedWin, _ := ConvertStringArrayToInt(strings.Split(guessedAndWinningNumbers[0], " "))
	parsedGuess, _ := ConvertStringArrayToInt(strings.Split(guessedAndWinningNumbers[1], " "))

	fmt.Println("Converting")
	c.winningNumbers = parsedWin
	c.guessedNumbers = parsedGuess

}

func CreateNewCard(input string) Card {

	newCard := Card{
		input:          input,
		winningNumbers: make([]int, 0),
		guessedNumbers: make([]int, 0),
	}

	newCard.ParseInput()

	return newCard
}

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	scanner.Scan()
	line := scanner.Text()
	card := CreateNewCard(line)

	fmt.Println(card.guessedNumbers)

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
