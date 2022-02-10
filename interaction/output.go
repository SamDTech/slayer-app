package interaction

import "fmt"

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")

}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please Choose your action")
	fmt.Println("---------------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player Attack Monster for %v damage\n", roundData.PlayerAttackDmg)

	} else if roundData.Action == "SPECIAL_ATTACT" {
		fmt.Printf("Player Performs a Strong Attack for %v damage\n", roundData.PlayerAttackDmg)

	} else {
		fmt.Printf("Player Heals for %v health\n", roundData.PlayerHealValue)
	}


	fmt.Printf("Monster Attack Player for %v damage\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)
}


func DeclareWinner(winner string) {
	fmt.Println("---------------------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("---------------------------------")

	fmt.Printf("%v won!\n", winner)
}
