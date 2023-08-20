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

		if strings.Contains(line, "1") {
			for i := 0; i < len(initialBlock); i++ {
				fmt.Println(initialBlock[i])
			}

			fmt.Println("----------------")

			initialBlock = transposeBlocks(initialBlock)
			break
		}

		row := strings.Split(line, " ")
		initialBlock = append(initialBlock, row)

	}

	for i := 0; i < len(initialBlock); i++ {
		fmt.Println(initialBlock[:][i])
	}

	//for sc.Scan() {
	//instructions := sc.Text()
	//move, from, to := parseInstructions(instructions)

	//var i int64
	//for i = 0; i < move; i++ {
	//TODO do the operations, you may need to transpose slices

	//}

	//}

	// for i := 0; i < len(initialBlock); i++ {
	// 	fmt.Println(initialBlock[:][i])
	// }

}

func transposeBlocks(slice [][]string) [][]string {
	var column []string
	var transposedBlocks [][]string = make([][]string, 0)

	for i := 0; i < 9; i++ {
		column = buildColumn(slice, i)
		transposedBlocks = append(transposedBlocks, column)
	}

	return transposedBlocks
}

func buildColumn(slice [][]string, columnIndex int) (column []string) {
	// Mixes column index 2 with index 5
	column = make([]string, 0)

	for _, row := range slice {
		fmt.Println(row[columnIndex])
		column = append(column, row[columnIndex])

	}
	fmt.Println()

	return column
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
