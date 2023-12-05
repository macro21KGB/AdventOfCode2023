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

type TransformationMapValue struct {
	SourceStartRange      int
	DestinationStartRange int
	Length                int
}

type ConversionMapType = map[string]map[int]int

// Extract the seeds from the input string
func ExtractSeeds(input string) []int {
	seedsString := strings.Split(input, ":")[1]
	seedsString = strings.TrimSpace(seedsString)

	return StringsToInts(strings.Split(seedsString, " "))
}

func ExtractFromTo(input []string) (string, string) {
	return input[1], input[2]
}

func UpdateConversionMap(conversionMap *ConversionMapType, from int, to int, where string) {
	if (*conversionMap)[where] == nil {
		(*conversionMap)[where] = make(map[int]int)
	}

	(*conversionMap)[where][from] = to

}

func FindSeedToLocationMap(conversionMap ConversionMapType, seed int, conversionSequence []string) int {

	currentConverted := seed
	for _, conversion := range conversionSequence {
		if conversionMap[conversion][currentConverted] != 0 {
			currentConverted = conversionMap[conversion][currentConverted]
		}
	}

	return currentConverted
}

func CalculateTransformationMap() (ConversionMapType, []int) {
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
	conversionMap := make(ConversionMapType, 0)

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
					SourceStartRange:      values[1],
					DestinationStartRange: values[0],
					Length:                values[2],
				})
			}

			transformationMaps = append(transformationMaps, newTMap)
		}

	}

	for _, tMap := range transformationMaps {

		for _, value := range tMap.Values {
			for i := 0; i < value.Length; i++ {
				UpdateConversionMap(&conversionMap, value.SourceStartRange+i, value.DestinationStartRange+i, tMap.To)
			}
		}

		log.Println(len(conversionMap))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return conversionMap, seeds
}

func main() {

	conversionMap, seeds := CalculateTransformationMap()
	for _, seed := range seeds {
		founded := FindSeedToLocationMap(conversionMap, seed, []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"})

		fmt.Println(founded)
	}

}
