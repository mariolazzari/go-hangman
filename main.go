package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// start the main program right here
func main() {
	// how many games won and lost
	wins := 0
	loses := 0

	// how long shall our word be
	numberOfLetters := 0

	// generate a random number from 4 to 15 and start a first round
	reInitRandomValue(&numberOfLetters)
	again, hasWon := playHangman(numberOfLetters)

	// start an quasi infinite loop
	for {
		if hasWon == true {
			wins++
		} else {
			loses++
		}

		// will we play again at all?
		if again == "y" {
			printScore(wins, loses)
			reInitRandomValue(&numberOfLetters)
			again, hasWon = playHangman(numberOfLetters)
		} else if again == "n" {
			printScore(wins, loses)
			break
		}
	}
	fmt.Println("Bye.")
}

// we reinitialize with a new random number between 4 and 15
func reInitRandomValue(toInit *int) {
	*toInit = rand.Intn(11) + 4
}

func chooseRandomWord(numberOfLetters int, gameType string) string {
	switch gameType {
	case "e":
		var lettersData []byte
		var err error
		if numberOfLetters == 4 {
			lettersData, err = os.ReadFile("words/simple4letters.txt")
		} else if numberOfLetters == 5 {
			lettersData, err = os.ReadFile("words/simple5letters.txt")
		} else if numberOfLetters >= 6 {
			lettersData, err = os.ReadFile("words/simple6letters.txt")
		}
		if err != nil {
			panic(err)
		}
		dataString := string(lettersData)
		someWords := strings.Split(dataString, " ")
		randomNumber := rand.Intn(len(someWords) - 1)
		chosenWord := someWords[randomNumber]
		return chosenWord

	case "h":
		var lettersData []byte
		var err error
		switch numberOfLetters {
		case 4:
			lettersData, err = os.ReadFile("words/all4letters.txt")
		case 5:
			lettersData, err = os.ReadFile("words/all5letters.txt")
		case 6:
			lettersData, err = os.ReadFile("words/all6letters.txt")
		case 7:
			lettersData, err = os.ReadFile("words/all7letters.txt")
		case 8:
			lettersData, err = os.ReadFile("words/all8letters.txt")
		case 9:
			lettersData, err = os.ReadFile("words/all9letters.txt")
		case 10:
			lettersData, err = os.ReadFile("words/all10letters.txt")
		case 11:
			lettersData, err = os.ReadFile("words/all11letters.txt")
		case 12:
			lettersData, err = os.ReadFile("words/all12letters.txt")
		case 13:
			lettersData, err = os.ReadFile("words/all13letters.txt")
		case 14:
			lettersData, err = os.ReadFile("words/all14letters.txt")
		case 15:
			lettersData, err = os.ReadFile("words/all15letters.txt")
		}
		if err != nil {
			panic(err)
		}
		dataString := string(lettersData)
		someWords := strings.Split(dataString, " ")
		randomNumber := rand.Intn(len(someWords) - 1)
		chosenWord := someWords[randomNumber]
		return chosenWord
	}
	return "this you will only see in case of a weird bug"
}
