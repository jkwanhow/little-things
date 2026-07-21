package main

import (
	"bufio"
	"log"
	"os"
)

/*
File used to format text to be ready to display for CLI.
*/

func CreateTargetText() string {
	targetText := ""

	file, err := os.Open("./text1.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		char := scanner.Text()
		targetText += char
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %s", err)
	}

	return targetText
}

func CreateTargetWords() []string {
	var targetWords []string

	file, err := os.Open("./text1.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words := scanner.Text()
		targetWords = append(targetWords, words)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %s", err)
	}

	return targetWords
}

func isEndOfString(word string, index int) bool {
	return len(word)-1 == index
}
