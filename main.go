package main

import (
	"fmt"

	"github.com/samdtech/slayer-app/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		executeRounds()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRounds() string {
	currentRound++
	isSpecialRound := currentRound % 3 == 0

	fmt.Print(isSpecialRound)

	interaction.ShowAvailableActions(isSpecialRound)

	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	if userChoice == "ATTACK" {

	}else if userChoice == "HEAL"{

	} else {

	}

	return ""

}

func endGame() {}
