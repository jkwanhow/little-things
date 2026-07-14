package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const LIVES = 6

const WRONG = "🟥"
const CORRECT = "🟩"
const HAS = "🟨"
const DEFAULT = "⬜"

// TODO: Fix the bug/issue where HAS shows up when the
// correct is already solved.
// e.g. steaa for steak gives GREENGREENGREENGREENORANGE
// should be GREENGREENGREENGREENRED
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
	cleaned := strings.ReplaceAll(s, "\n", "")
	cleaned = strings.ToLower(cleaned)
	return cleaned
}

func CreateDictionary() map[string]bool {
	dict := map[string]bool{}
	// 5 letter words
	file, err := os.Open("./valid_words.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		dict[word] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %s", err)
	}

	return dict
}

func main() {
	const answer = "grind"
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
