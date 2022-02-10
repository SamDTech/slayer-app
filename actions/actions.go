package actions

import (
	"math/rand"
	"time"
)

// generate rand source
var randSource = rand.NewSource(time.Now().UnixNano())

var currentMonsterHealth = 100
var currentPlayerHealth = 100

// generate random number
var randGenerator = rand.New(randSource)

func AttackMoster(isSpecialAttack bool) {
	minAttackValue := 5
	maxAttackValue := 10

	if isSpecialAttack {
		minAttackValue = 10
		maxAttackValue = 20
	}

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	// subtract monster health from the damage value
	currentMonsterHealth -= dmgValue
}

func HealPlayer() {
	minHealValue := 10
	maxHealValue := 20

	healValue := generateRandBetween(minHealValue, maxHealValue)

	healthDiff := 100 - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = 100
	}

	// add player health from the damage value
	currentPlayerHealth += healValue
}

func AttackPlayer() {
	minAttackValue := 9
	maxAttackValue := 13

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	// subtract monster health from the damage value
	currentPlayerHealth -= dmgValue
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
