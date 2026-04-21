package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// clear the screen depending on your OS
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

// we print the actual score only
func printScore(wins, loses int) {
	clearScreen()
	fmt.Printf("#####################\n")
	fmt.Printf("     Your Score\n")
	fmt.Printf(" %d: wins, %d: loses\n", wins, loses)
	fmt.Printf("#####################\n")
	fmt.Println()
}

// create a string of the length of a word as a String of underlines
func hideTheWord(wordLength int) string {
	var dashes strings.Builder
	for range wordLength {
		dashes.WriteString("_")
	}
	return dashes.String()
}

// exchanges an underline with a letter in case it was guessed right
func revealDashes(word string, guess string, dashes string) string {
	var newDashes strings.Builder
	for i, r := range dashes {
		if c := string(r); c != "_" {
			newDashes.WriteString(c)
		} else {
			var letter = string(word[i])
			if guess == letter {
				newDashes.WriteString(strings.ToUpper(guess))
			} else {
				newDashes.WriteString("_")
			}
		}
	}
	return newDashes.String()
}
