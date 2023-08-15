package main

import (
	"fmt"
	"bufio"
	"os"
)

type StringSet map[rune]struct{}

func handleError (err error) {
	if err != nil {
		panic(err)
	}
}

func main () {
	file, err := os.Open("./input.txt")	
	handleError(err)
	defer file.Close()
	sc := bufio.NewScanner(file)

	sum := 0

	for sc.Scan() {
		line := sc.Text()
		ruckSackPriority := findSharedItemTypesAndCalculateRucksackPriority(line)
		sum = sum + ruckSackPriority
	}

	fmt.Printf(" %d \n", sum)
}

func findSharedItemTypesAndCalculateRucksackPriority(rucksack string) int {
	halfOfLen := len(rucksack) /2
	firstCompartment := rucksack[0:halfOfLen]
	secondCompartment := rucksack[halfOfLen:]
	commonLetters := intersectCompartments(firstCompartment, secondCompartment)
	ruckSackPriority := calculateRuckSackPriority(commonLetters)

	return ruckSackPriority	
}

func intersectCompartments(firstCompartment, secondCompartment string) StringSet {
	mapOfFirst := make(map[rune]int)
	commonLetters := make(StringSet)

	for _, letter := range firstCompartment {
		mapOfFirst[letter]++
	}

	for _, letter := range secondCompartment {
		_, isRecorded := commonLetters[letter]
		if mapOfFirst[letter] > 0 && !isRecorded {
			commonLetters[letter] = struct{}{}
		}
	}

	return commonLetters
}

func calculateRuckSackPriority(commonLetters StringSet) int {
	sum := 0
	for k, _ := range commonLetters {
		letterPriority := calculateLetterPriority(k)
		sum = sum + letterPriority
	}

	return sum
}

func calculateLetterPriority(letter rune) int {
	ascii := int(letter)
	var priority int = 0
	// Is key in Capital
	if ascii >= 65 && ascii <= 90 {
		priority = ascii - 65 + 27
	} else if ascii >= 97 && ascii <= 122 {
		priority = ascii - 97 +1
	}
	
	return priority
}
