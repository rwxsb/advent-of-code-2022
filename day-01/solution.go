package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var topMostCalories int64 = 0
	var secondMostCalories int64 = 0
	var thirdMostCalories int64 = 0
	var sum int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > topMostCalories {
				if topMostCalories > secondMostCalories {
					if secondMostCalories > thirdMostCalories {
						thirdMostCalories = secondMostCalories
					}
					secondMostCalories = topMostCalories
				} else if topMostCalories > thirdMostCalories {
					thirdMostCalories = topMostCalories
				}
				topMostCalories = sum
			} else if sum > secondMostCalories {
				if secondMostCalories > thirdMostCalories {
					thirdMostCalories = secondMostCalories
				}
				secondMostCalories = sum
			} else if sum > thirdMostCalories {
				thirdMostCalories = sum
			}
			sum = 0
			continue
		}

		lineAsInt, err := strconv.ParseInt(line, 10, strconv.IntSize)
		handleError(err)
		sum = sum + lineAsInt
	}

	fmt.Printf("1st %d \n", topMostCalories)
	fmt.Printf("2nd %d \n", secondMostCalories)
	fmt.Printf("3rd %d \n", thirdMostCalories)
	fmt.Printf("sum %d \n", topMostCalories+secondMostCalories+thirdMostCalories)

}
