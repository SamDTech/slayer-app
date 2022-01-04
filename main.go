package main
import "github.com/samdtech/slayer-app/interaction"


var currentRound = 0


func main(){
	startGame()

	winner := ""

	for winner == "" {
		executeRounds()
	}

	endGame()
}

func startGame(){
	interaction.PrintGreeting()
}

func executeRounds() {
	currentRound++
	isSpecialRound := currentRound % 3 == 0
}

func endGame() { }