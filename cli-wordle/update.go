package main

import (
	"fmt"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
)

// All functions are abstractions of update logic
func (m *model) EnterLine() {
	splitGuess := m.guesses[m.curRow]
	var guess string
	for _, char := range splitGuess {
		guess += char
	}
	guess = CleanString(guess)
	if len(guess) != 5 {
		m.message = "Not enough letters"
		m.state = "warn"

	} else if !m.dictionary[guess] {
		m.message = "Not in word list"
		m.state = "warn"
	} else {
		// need to check if the word is correct too.
		m.letterStates[m.curRow] = GetStatesOfLetters(m.answer, guess)
		if guess == m.answer {
			m.state = "end"
			m.message = fmt.Sprintf("Nice, you got it in %d turns", m.curRow+1)
		} else if m.curRow == 5 {
			m.state = "end"
			m.message = fmt.Sprintf("[%s]", m.answer)
		}
		// process the colors and all
		m.curRow++
		m.curCol = 0
	}
}

func (m *model) CharacterInput(msg tea.KeyMsg) {
	m.message = ""
	m.state = "play"
	if m.curCol < 5 {
		keyStr := msg.String()
		if len(keyStr) == 1 && unicode.IsLetter(rune(keyStr[0])) {
			letter := keyStr
			m.guesses[m.curRow][m.curCol] = CleanString(letter)
			m.curCol++
		}
	}
}

func (m *model) HandleDelete() {
	m.message = ""
	m.state = "play"
	if m.curCol > 0 {
		m.curCol--
		m.guesses[m.curRow][m.curCol] = ""
	}
}
