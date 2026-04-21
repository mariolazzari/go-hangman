package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

// init random generation
// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func main() {
	wins := 0
	loses := 0

	// word len
	numberOfLetters := 0

	// generate a random letter and start 1st round
	reInitRandonValue(&numberOfLetters)
	again, hasWon := playHangman(numberOfLetters)

	for {
		if hasWon {
			wins++
		} else {
			loses++
		}

		// play again?
		if again == "y" {
			printScore(wins, loses)
			reInitRandonValue(&numberOfLetters)
			//again, hasWon := playHangman(numberOfLetters)

		} else if again == "n" {
			printScore(wins, loses)
			break
		}

	}
}

func reInitRandonValue(toInit *int) {
	*toInit = rand.Intn(11) + 4
}

func playHangman(n int) (string, bool) {
	return "", false
}

func printScore(w, l int) {
	clearScreen()
	fmt.Printf("Yor score: %d-%d", w, l)
}

func clearScreen() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
