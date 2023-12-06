package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	t "time"
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

func JoinInts(ints []int, sep string) string {
	strs := make([]string, 0, len(ints))
	for _, i := range ints {
		strs = append(strs, strconv.Itoa(i))
	}
	return strings.Join(strs, sep)
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

	time, _ := strconv.Atoi(JoinInts(times, ""))
	distance, _ := strconv.Atoi(JoinInts(distances, ""))

	startTime := t.Now()
	log.Println("Start calculating at ", startTime.Format("15:04:05.0000"))
	final := CalculateNumberOfWins(time, distance)

	log.Printf("Calculation took %s", t.Since(startTime))
	log.Println("Final result: ", final)
}
