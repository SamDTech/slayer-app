package main

import (
	"github.com/samdtech/slayer-app/actions"
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
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)

	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerHeath int
	var monsterHealth int

	if userChoice == "ATTACK" {
		actions.AttackMoster(false)

	} else if userChoice == "HEAL" {
		 actions.HealPlayer()
	} else {
		// now we are in special round
		 actions.AttackMoster(true)
	}

	// monster should always attack
	 actions.AttackPlayer()

	// get the current health
	playerHeath, monsterHealth = actions.GetCurrentPlayerHealth()

	// check if player or monster is dead
	if playerHeath <= 0 {
		return "Monster"
	} else if monsterHealth <= 0{
		return "Player"
	}

	return ""

}

func endGame() {}
