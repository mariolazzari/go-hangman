package main

import (
	"fmt"
	"regexp"
	"strings"
)

// play a round of hangman
func playHangman(numberOfLetters int) (playAgain string, isWinner bool) {
	stageOfHangman := 0        // how far is hangman from death
	gameType := ""             // easy or hard
	hasGuessedALetter := false // do we have a hit?
	hasWon := false            // is the actual game won or not?
	guess := ""                // actual letter guessed
	guessedLetters := ""       // a string of guessed letters
	again := ""                // type "y" or "Y" for another round
	dashes := ""               // word to guess completely hidden in dashes
	newDashes := ""            // word partly relvealed
	fmt.Println("#####################")
	fmt.Printf("### H A N G M A N ###\n")
	fmt.Println("#####################")
	fmt.Println()
	for {
		fmt.Println("Select game type:")
		fmt.Println("[e] Easy - use simple 4,5 and 6 letter words only")
		fmt.Println("[h] Hard - use more difficult words with 4 to 15 letters")
		fmt.Scanln(&gameType)
		// no need to check uppercase like
		// (gameType == "e") || (gameType == "E") || (gameType =="h") || (gameType =="H")
		// needed if we just convert all to lowercase
		gameType = strings.ToLower(gameType)
		if (gameType == "e") || (gameType == "h") {
			clearScreen()
			break
		} else {
			fmt.Println("Please press [e] or [h] key")
		}
	}

	// choose a word to guess
	word := chooseRandomWord(numberOfLetters, gameType)

	// start the main loop here
	fmt.Println()
	for {
		drawHangman(stageOfHangman)
		if stageOfHangman == 10 {
			fmt.Printf("OMG, you let hangman die!\n")
			fmt.Printf("You could have saved him with the word: %s\n", strings.ToUpper(word))
			return wantToPlayAgain(), false
		}

		if hasGuessedALetter == false {
			dashes = hideTheWord(len(word))
			fmt.Printf(" %s\n", dashes)
		} else {
			fmt.Printf(" %s\n", newDashes)
		}
		fmt.Printf("\n Guess a letter: ")
		fmt.Scanln(&guess)

		isALetter, someKindOfError := regexp.MatchString("^[a-zA-Z]", guess)
		if someKindOfError != nil {
			clearScreen()
			fmt.Printf("Something went horribly wrong. ")
			fmt.Printf("Exiting with error can not regex match %v", again)
			return
		}

		// we allow upper case but continue with lower case only
		guess = strings.ToLower(guess)

		if !isALetter {
			clearScreen()
			fmt.Printf("That is not even a letter. Try again.\n")
		} else if len(guess) > 1 {
			clearScreen()
			fmt.Printf("You entered more than 1 character. Try again.\n")
		} else if strings.Contains(guessedLetters, guess) {
			clearScreen()
			fmt.Printf("You have already guessed that letter. Try again.\n")
		} else if strings.Contains(word, guess) {
			clearScreen()
			fmt.Printf("Great! The letter you guessed is in the word.\n")
			guessedLetters += guess

			if !hasGuessedALetter {
				updatedDashes := revealDashes(word, guess, dashes)
				newDashes = updatedDashes
			} else {
				updatedDashes := revealDashes(word, guess, newDashes)
				newDashes = updatedDashes
			}

			// we have a hit, a good time to check
			// if all revealed letters are identical
			// if yes the player won
			hasGuessedALetter = true
			if newDashes == strings.ToUpper(word) {
				hasWon = true
			}

			// output in case of a winning a round
			if hasWon == true {
				clearScreen()
				fmt.Printf("### C O N G R A T U L A T I O N S ###\n")
				fmt.Println()
				fmt.Printf(" _   _\n")
				fmt.Printf("  \\O/\n")
				fmt.Printf("   |\n")
				fmt.Printf("   |\n")
				fmt.Printf("  / \\\n")
				fmt.Println()
				fmt.Printf("You won this round! The word was: %s\n", strings.ToUpper(word))
				fmt.Printf("You saved hangman in %v of 10 guesses.\n", stageOfHangman)
				return wantToPlayAgain(), true
			}
		} else {
			clearScreen()
			fmt.Printf("The letter you guessed is not in the word. :(\n")
			stageOfHangman++
			guessedLetters += guess
		}
	}
}

func wantToPlayAgain() string {
	for {
		again := ""
		fmt.Printf("Play again? [y/n]\n")
		fmt.Scanln(&again)
		isYorN, someKindOfError := regexp.MatchString("^y|Y|n|N", again)
		if someKindOfError != nil {
			fmt.Printf("Something went horribly wrong. ")
			fmt.Printf("Exiting with error can not regex match %v", again)
			panic(someKindOfError)
		}
		if !isYorN {
			fmt.Printf("You didn't type [y] or [n}. Try again.\n")
		} else if len(again) > 1 {
			fmt.Printf("You entered more than 1 character. Try again.\n")
		} else if strings.ToLower(again) == "y" {
			return "y"
		} else {
			return "n"
		}
	}
}

// we draw hangman depending on his actual state (number of wrong guesses)
func drawHangman(stageOfHangman int) {
	switch stageOfHangman {
	case 0:
		fmt.Printf("   +---+\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 1:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 2:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 3:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf("  /|   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 4:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 5:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\  |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 6:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\_ |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 7:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\_ |\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 8:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\_ |\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("  /    |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 9:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   O   |\n")
		fmt.Printf(" _/|\\_ |\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("  / \\  |\n")
		fmt.Printf("       |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	case 10:
		fmt.Printf("   +---+\n")
		fmt.Printf("   |   |\n")
		fmt.Printf("   0   |\n")
		fmt.Printf("  /|\\  |\n")
		fmt.Printf(" ° | ° |\n")
		fmt.Printf("  / \\  |\n")
		fmt.Printf("       |\n")
		fmt.Printf(" R.I.P.|\n")
		fmt.Printf(" ======== %v/10 Guesses\n", stageOfHangman)
		fmt.Println()

	}
}
