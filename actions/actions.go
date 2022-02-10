package actions

import (
	"math/rand"
	"time"
)

// generate rand source
var randSource = rand.NewSource(time.Now().UnixNano())

var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

// generate random number
var randGenerator = rand.New(randSource)


// function to Attack Monster
func AttackMoster(isSpecialAttack bool)  {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	// subtract monster health from the damage value
	currentMonsterHealth -= dmgValue

}

func HealPlayer()  {

	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)

	healthDiff := PLAYER_HEALTH - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
	}

	// add player health from the damage value
	currentPlayerHealth += healValue


}

func AttackPlayer()  {
	minAttackValue := MONSTER_ATTACK_MIN_DMG
	maxAttackValue := MONSTER_ATTACK_MAX_DMG

	// get the damage value
	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	// subtract monster health from the damage value
	currentPlayerHealth -= dmgValue


}


// function to get current player and monster health
func GetCurrentPlayerHealth() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

// function to generate random number between min and max
func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
