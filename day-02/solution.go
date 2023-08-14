package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rules struct {
	beats string
}

func handleError (err error) {
	if err != nil {
		panic(err)
	}
}

func main () {
	scoreOfCard := make(map[string]int)
	scoreOfCard["Rock"] = 1 
	scoreOfCard["Paper"] = 2 
	scoreOfCard["Scissors"] = 3

	letterToHandMap := make(map[string]string)

	letterToHandMap["X"] = "Rock"
	letterToHandMap["Y"] = "Paper"
	letterToHandMap["Z"] = "Scissors"
	letterToHandMap["A"] = "Rock"
	letterToHandMap["B"] = "Paper"
	letterToHandMap["C"] = "Scissors"

	file, err := os.Open("./input.txt")	
	handleError(err)
	sc := bufio.NewScanner(file)

	sum := 0
	for sc.Scan() {
		line := sc.Text()
		round := strings.Split(line, " ")
		opponnetHand := letterToHandMap[round[0]]
		myHand := letterToHandMap[round[1]]		
		roundResult := roundScoreCheat([2]string{opponnetHand, myHand})
		extraPoints := scoreOfCard[myHand]
		sum = sum + roundResult + extraPoints
	}

	fmt.Printf("%d \n",sum)

}

func roundScore (round [2]string) int {
	ruleSet := make(map[string]Rules)

	ruleSet["Rock"] = Rules{ beats: "Scissors"}
	ruleSet["Paper"] = Rules{ beats: "Rock"} 
	ruleSet["Scissors"] = Rules{ beats: "Paper"}


	if round[0] == round[1] {
		return 3
	}

	if ruleSet[round[1]].beats == round[0] {
		return 6
	}

	return 0

}


func roundScoreCheat (round [2]string) int {
	ruleSet := make(map[string]Rules)

	ruleSet["Rock"] = Rules{ beats: "Scissors"}
	ruleSet["Paper"] = Rules{ beats: "Rock"} 
	ruleSet["Scissors"] = Rules{ beats: "Paper"}


	if round[0] == "Scissor" {
		return 6
	}

	if round[0] == "Paper" {
		return 3
	}

	return 0

}    


