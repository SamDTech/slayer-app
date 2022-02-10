package main

import (
	"github.com/samdtech/slayer-app/actions"
	"github.com/samdtech/slayer-app/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := "" // "Player" or "Monster" or ""

	for winner == "" {
		winner = executeRounds()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRounds() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)

	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int // default value = 0
	var monsterAttackDmg int
	var playerHealValue int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMoster(false)

	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		// now we are in special round
		playerAttackDmg = actions.AttackMoster(true)
	}

	// monster should always attack
	monsterAttackDmg = actions.AttackPlayer()

	// get the current health
	playerHeath, monsterHealth := actions.GetCurrentPlayerHealth()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
		PlayerHealth:     playerHeath,
		MonsterHealth:    monsterHealth,
	}

	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	// check if player or monster is dead
	if playerHeath <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""

}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
