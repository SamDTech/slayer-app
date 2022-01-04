package main
import "github.com/samdtech/slayer-app/interaction"



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

func executeRounds() {}

func endGame() { }