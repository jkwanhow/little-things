package main

import (
	"bufio"
	"fmt"
	"os"
)

const LIVES = 6

const WRONG = "🟥"
const CORRECT = "🟩"
const HAS = "🟨"
const DEFAULT = "⬜"

func GetNotCorrectSquare(c rune, a string) string {
	for _, aChar := range a {
		if aChar == c {
			return HAS
		}
	}

	return WRONG
}

func CreateSquareOutput(a string, g string) string {
	// remember g and output share the same positioning
	// in terms of rune/char to square
	output := make([]string, len(a))

	workingCopy := []rune(a)
	for pos, char := range g {
		if char == rune(a[pos]) {
			workingCopy[pos] = '_'
			output[pos] = CORRECT

		}
	}

	for pos, char := range g {
		if output[pos] != CORRECT {
			output[pos] = GetNotCorrectSquare(char, string(workingCopy))
		}
	}

	var joinedOutput string
	for _, char := range output {
		joinedOutput += char
	}

	return joinedOutput
}

func GetLength(s string) int {
	var length int
	for range s {
		length++
	}
	return length
}

func old_main() {
	const answer = "pshaw"
	length := GetLength(answer)
	reader := bufio.NewReader(os.Stdin)
	dictionary := CreateDictionary()

	attempt := 0
	fmt.Print("Go Wordle! \n")
	for attempt < 6 {
		guess, _ := reader.ReadString('\n')
		guess = CleanString(guess)
		if GetLength(guess) != length {
			fmt.Printf("Keep the word %d character\n", length)
		} else if !dictionary[guess] {
			fmt.Print("Word not found in dictionary\n")
		} else {
			fmt.Println(CreateSquareOutput(answer, guess))
			if guess == answer {
				break
			}
			attempt += 1
		}
	}

	if attempt > 6 {
		fmt.Printf("Didn't get the word")
	} else {
		fmt.Printf("Nice")
	}

}
