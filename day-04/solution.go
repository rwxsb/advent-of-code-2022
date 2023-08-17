package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func handleError (err error) {
	if err != nil {
		panic(err)
	}
}

func main () {


	file, err := os.Open("input.txt")
	handleError(err)
	sc := bufio.NewScanner(file)
	count := 0

	for sc.Scan() {
		line := sc.Text()
		assingments := strings.Split(line, ",")
		count = count + compareAssignments(assingments[0],assingments[1])
	}

	fmt.Println(count)
}


func compareAssignments(first, second string) int {
	firstAssingment := strings.Split(first, "-")
	secondAssingment := strings.Split(second, "-")

	firstStart, _ := strconv.ParseInt(firstAssingment[0],10,64)
	firstEnd, _ := strconv.ParseInt(firstAssingment[1],10, 64)

	secondStart, _ := strconv.ParseInt(secondAssingment[0],10, 64)
	secondEnd, _ := strconv.ParseInt(secondAssingment[1],10, 64)

	if firstStart <= secondStart && secondEnd <= firstEnd {
		return 1
	} else if secondStart <= firstStart && firstEnd <= secondEnd {
		return 1
	}

	return 0
}
