package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const (
	PlayerChoiceAttack = iota + 1
	PlayerChoiceHeal
	PlayerChoiceSpecialAttack
)

func GetPlayerChoice(specialAttackIsAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()
		if playerChoice == fmt.Sprint(PlayerChoiceAttack) {
			return "ATTACK"
		} else if playerChoice == fmt.Sprint(PlayerChoiceHeal) {
			return "HEAL"
		} else if playerChoice == fmt.Sprint(PlayerChoiceSpecialAttack) && specialAttackIsAvailable {
			return "SPECIAL_ATTACK"
		}

		fmt.Println("Fetching the user input failed. Please try again!")

	}

}

func getPlayerInput() (string, error) {
	fmt.Print("Your Choice:  ")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput, nil

}