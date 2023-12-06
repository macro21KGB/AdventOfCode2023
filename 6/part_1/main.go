package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ExtractNumbers(str string) []int {
	re := regexp.MustCompile(`\d+`)
	numStrs := re.FindAllString(str, -1)
	numbers := make([]int, 0, len(numStrs))

	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}

	return numbers
}

type Game struct {
	MaxTime     int
	MinDistance int
}

func CalculateNumberOfWins(maxTime int, minDistance int) int {

	wins := 0
	for i := 1; i < maxTime; i++ {
		howManyMillisRemaining := maxTime - i
		distanceTraversed := howManyMillisRemaining * i

		if distanceTraversed > minDistance {
			wins++
		}

	}

	return wins
}

func main() {

	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	times := make([]int, 0)
	distances := make([]int, 0)
	isTimes := true
	for scanner.Scan() {
		line := scanner.Text()
		if isTimes {
			times = ExtractNumbers(line)
			isTimes = false
		} else {
			distances = ExtractNumbers(line)
		}
	}

	totalPower := 1
	for i := 0; i < len(distances); i++ {
		totalPower *= CalculateNumberOfWins(times[i], distances[i])
	}

	log.Println(totalPower)
}
