package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)

	asciiFigure.Print()

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
	asciiFigure := figure.NewColorFigure("GAME OVER!", "","red", true)

	asciiFigure.Print()
	fmt.Println("---------------------------------")

	fmt.Printf("%v won!\n", winner)
}


func WriteLogFile(roundData *[]RoundData) {
	exPath, err := os.Executable()



	if err != nil {
		fmt.Println("Writing log file failed. Exiting...")
		return
	}

	exPath = filepath.Dir(exPath)


	// create a file
	file, err := os.Create(exPath + "/gameLog.txt")

	// for go run .
	//file, err := os.Create("gameLog.txt")

	if err != nil {
		fmt.Println("Saving a log file failed. Exiting", err)
		return
	}

	// write data into the file
	for index, value := range *roundData{
		logEntry := map[string]string{
			"Round": fmt.Sprint(index + 1),
			"Action": value.Action,
			"Player Attack Damage": fmt.Sprint(value.PlayerAttackDmg),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Heal Value": fmt.Sprint(value.PlayerHealValue),
			"Player Health": fmt.Sprint(value.PlayerHealth),
			"Monster Health": fmt.Sprint(value.MonsterHealth),
		}

		logLine := fmt.Sprintln(logEntry)

		// write the line to the file
	_, err :=	file.WriteString(logLine)

	if err != nil {
		fmt.Println("Writing into log file failed. Exiting")
		continue

	}

	}

	// close the file
	file.Close()

	fmt.Println("Wrote data to log.")

}