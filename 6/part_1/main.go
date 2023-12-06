package main

import (
	"bufio"
	"fmt"
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

func CalculateProductNumberOfWins(maxTime int, minDistance int) int {

	result := 1
	for i := 1; i < maxTime; i++ {
		howManyMillisRemaining := maxTime - i
		distanceTraversed := howManyMillisRemaining * i

		if distanceTraversed > minDistance {
			result *= distanceTraversed
		}

	}

	return result
}

func main() {

	file, _ := os.Open("./test_input.txt")
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

	totalSum := 0
	for i := 0; i < len(distances); i++ {
		resultProduct := CalculateProductNumberOfWins(times[i], distances[i])
		fmt.Println(resultProduct)
	}

	log.Println(totalSum)
}
