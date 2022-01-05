package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")

}

func ShowAvailableActions(specialAttackIsAvailable bool){
	fmt.Println("Please Choose your action")
	fmt.Println("---------------------------------")
}