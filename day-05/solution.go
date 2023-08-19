package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	file, err := os.Open("input.txt")
	handleError(err)
	sc := bufio.NewScanner(file)
	var initialBlock [][]string

	for sc.Scan() {
		line := sc.Text()
		row := strings.Split(line, " ")
		initialBlock = append(initialBlock, row)

		if strings.Contains(line, " 1   2   3   4   5   6   7   8   9 ") {
			break
		}
	}

	for i := 0; i < len(initialBlock); i++ {
		fmt.Println(initialBlock[i][0])
	}

	fmt.Println("---------------------")

	for i := 0; i < len(initialBlock[0]); i++ {
		fmt.Println(initialBlock[0][i])
	}

	for sc.Scan() {
		instructions := sc.Text()
		move, from, to := parseInstructions(instructions)

		var i int64
		for i = 0; i < move; i++ {
			//TODO do the operations, you may need to transpose slices

		}

	}

	// for i := 0; i < len(initialBlock); i++ {
	// 	fmt.Println(initialBlock[:][i])
	// }

}

func parseInstructions(instructions string) (move, from, to int64) {
	wordSplit := strings.Split(instructions, " ")

	for i := 0; i < len(wordSplit); i++ {

		if wordSplit[i] == "move" {
			result, err := strconv.ParseInt(wordSplit[i+1], 10, 64)
			handleError(err)
			move = result

		}

		if wordSplit[i] == "from" {
			result, err := strconv.ParseInt(wordSplit[i+1], 10, 64)
			handleError(err)
			from = result

		}

		if wordSplit[i] == "to" {
			result, err := strconv.ParseInt(wordSplit[i+1], 10, 64)
			handleError(err)
			to = result
		}

	}

	return move, from, to

}
