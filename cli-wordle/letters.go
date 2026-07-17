package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const HEIGHT = 5
const EMPTYLINE = "     "
const SPACE = " "

func CreateFontMap() map[string][]string {
	letterFonts := map[string][]string{}

	file, err := os.Open("./fontedLetters.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var currentLetter string
	for scanner.Scan() {
		line := scanner.Text()
		if after, ok := strings.CutPrefix(line, "@"); ok {
			currentLetter = after
			letterFonts[currentLetter] = []string{}

		} else {
			letterFonts[currentLetter] = append(letterFonts[currentLetter], line)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %s", err)
	}

	return letterFonts
}
