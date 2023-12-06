package main

import (
	"regexp"
	"strconv"
)

// Convert a slice of strings to a slice of ints
func StringsToInts(input []string) []int {
	output := make([]int, len(input))
	for i, value := range input {
		output[i], _ = strconv.Atoi(value)
	}

	return output
}

type Range struct {
	Start int
	End   int
}

func SelectMinimumValue(values []int) int {
	minimum := values[0]
	for _, value := range values {
		if value < minimum {
			minimum = value
		}
	}

	return minimum
}

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
