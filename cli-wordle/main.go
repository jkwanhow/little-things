package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	output := ""
	for pos, char := range g {
		if char == rune(a[pos]) {
			output += CORRECT

		} else {
			output += GetNotCorrectSquare(char, a)
		}
	}

	return output
}

func GetLength(s string) int {
	var length int
	for range s {
		length++
	}
	return length
}

func CleanString(s string) string {
	return strings.ReplaceAll(s, "\n", "")
}

func main() {
	const answer = "write"
	length := GetLength(answer)
	reader := bufio.NewReader(os.Stdin)

	attempt := 0
	fmt.Print("Go Wordle! \n")
	for attempt < 6 {
		guess, _ := reader.ReadString('\n')
		guess = CleanString(guess)
		if GetLength(guess) != length {
			fmt.Printf("Keep the word %d character\n", length)
		} else {
			fmt.Println(CreateSquareOutput(answer, guess))
			attempt += 1
		}
	}

}
