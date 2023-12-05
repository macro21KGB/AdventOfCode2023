package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Convert a slice of strings to a slice of ints
func StringsToInts(input []string) []int {
	output := make([]int, len(input))
	for i, value := range input {
		output[i], _ = strconv.Atoi(value)
	}

	return output
}

type TransformationMap struct {
	From   string
	To     string
	Values []TransformationMapValue
}

type Range struct {
	Start int
	End   int
}

type TransformationMapValue struct {
	SourceRange      Range
	DestinationRange Range
}

// Extract the seeds from the input string
func ExtractSeeds(input string) []int {
	seedsString := strings.Split(input, ":")[1]
	seedsString = strings.TrimSpace(seedsString)

	return StringsToInts(strings.Split(seedsString, " "))
}

func ExtractFromTo(input []string) (string, string) {
	return input[1], input[2]
}

func (tMap TransformationMap) ConvertSourceToDestination(valueToConvert int) int {
	for _, value := range tMap.Values {
		if value.SourceRange.Start <= valueToConvert && valueToConvert <= value.SourceRange.End {
			return valueToConvert + (value.DestinationRange.Start - value.SourceRange.Start)
		}
	}

	return valueToConvert
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

func FindSeedToLocationMap(maps []TransformationMap, seed int) int {

	currentConverted := seed

	for _, translationMap := range maps {
		newValue := translationMap.ConvertSourceToDestination(currentConverted)
		if newValue != currentConverted {
			currentConverted = newValue
		}

	}

	return currentConverted
}

func CalculateTransformationMap() ([]TransformationMap, []int) {
	mapRegex := regexp.MustCompile(`(\w+)\-to\-(\w+) map:`)
	// First Star
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	startLine := true
	transformationMaps := make([]TransformationMap, 0)
	seeds := make([]int, 0)
	// conversionSequence := []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

	// conversionMap è una mappa di mappe. La chiave esterna è una stringa, che corrisponde al valore dove vuoi andare,
	// mentre la chiave interna e il valore sono entrambi interi e sono le conversioni, se non c'è una conversione vale lo stesso numero

	// Loop through the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if startLine {
			seeds = ExtractSeeds(line)
			startLine = false
		}

		// found a map
		if mapRegex.MatchString(line) {
			matched := mapRegex.FindAllStringSubmatch(line, -1)[0]
			from, to := ExtractFromTo(matched)
			newTMap := TransformationMap{
				From:   from,
				To:     to,
				Values: make([]TransformationMapValue, 0),
			}

			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					break
				}

				values := StringsToInts(strings.Split(strings.TrimSpace(line), " "))
				newTMap.Values = append(newTMap.Values, TransformationMapValue{
					SourceRange:      Range{Start: values[1], End: values[1] + values[2] - 1},
					DestinationRange: Range{Start: values[0], End: values[0] + values[2] - 1},
				})
			}

			transformationMaps = append(transformationMaps, newTMap)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return transformationMaps, seeds
}

func main() {

	transformationMaps, seeds := CalculateTransformationMap()

	locations := make([]int, 0)
	for _, seed := range seeds {
		locations = append(locations, FindSeedToLocationMap(transformationMaps, seed))
	}

	fmt.Println(SelectMinimumValue(locations))
}
