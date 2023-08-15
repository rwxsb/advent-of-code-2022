package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rules struct {
	beats string
	loses string
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

	resultMap := make(map[string]string)
	resultMap["X"] = "Lose"
	resultMap["Y"] = "Draw"
	resultMap["Z"] = "Win"

	file, err := os.Open("./input.txt")	
	handleError(err)
	defer file.Close()
	sc := bufio.NewScanner(file)

	sum := 0
	for sc.Scan() {
		line := sc.Text()
		round := strings.Split(line, " ")
		opponentHand := letterToHandMap[round[0]]
		//Part 1 
		//myHand := letterToHandMap[round[1]]
		//roundResult := roundScore([2]string{opponentHand, myHand})
		result := resultMap[round[1]]
		roundResult := fullRoundScore(opponentHand, result, scoreOfCard)
		sum = sum + roundResult //+ scoreOfCard[round[1]] 
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

func fullRoundScore (opponentHand string, result string, scoreOfCard map[string]int) int {
	ruleSet := make(map[string]Rules)

	ruleSet["Rock"] = Rules{ beats: "Scissors", loses: "Paper" }
	ruleSet["Paper"] = Rules{ beats: "Rock", loses: "Scissors" } 
	ruleSet["Scissors"] = Rules{ beats: "Paper", loses: "Rock" }

	// return extra points + win, lose

	myHand := findMyPlay(opponentHand, result, ruleSet)
	myHandScore := scoreOfCard[myHand] 

	if result == "Win" {
		return 6 + myHandScore; 
	}

	if result == "Draw" {
		return 3 + myHandScore
	}

	return myHandScore

}    

func findMyPlay(opponentHand string, result string, ruleSet map[string]Rules) string{
	if result == "Win" {
		return	ruleSet[opponentHand].loses
	}

	if result == "Lose" {
		return ruleSet[opponentHand].beats
	}

	return opponentHand
	
}


