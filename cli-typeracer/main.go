package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor         int
	targetChar     rune
	targetText     string
	inputChar      rune
	userTotalInput string
}

const (
	Reset  = "\033[0m"
	Green  = "\033[42;37m"
	Yellow = "\033[43;37m"
	Orange = "\033[48;5;208;37m"
	Gray   = "\033[100;37m"
)

func (m model) View() string {
	s := ""
	lineLimit := 60
	targetPtr := 0
	inputPtr := 0
	correct := 0

	for targetPtr < len(m.targetText)-1 {
		if targetPtr%lineLimit == 0 && targetPtr != 0 {
			s += "\n"
			for inputPtr < len(m.userTotalInput) && inputPtr < targetPtr {
				s += string(m.userTotalInput[inputPtr])
				inputPtr++
			}
			s += "\n\n"

		}
		spaceColor := Reset
		if targetPtr < len(m.userTotalInput) {
			if m.userTotalInput[targetPtr] == m.targetText[targetPtr] {
				spaceColor = Green
				correct++
			} else {
				spaceColor = Orange
			}
		}
		s += fmt.Sprintf("%s%s%s", spaceColor, string(m.targetText[targetPtr]), Reset)
		targetPtr++
	}
	// last run to clear last line
	s += "\n"
	for inputPtr < len(m.userTotalInput) && inputPtr < targetPtr {
		s += string(m.userTotalInput[inputPtr])
		inputPtr++
	}

	accuracy := 100.00
	if inputPtr > 0 {
		accuracy = float64(correct) / float64(inputPtr) * 100
	}
	s += fmt.Sprintf("Accuracy: %.2f%%", accuracy)
	s += fmt.Sprintf("\n%s\n", string(m.inputChar))
	return s
}

func initialModel() model {
	targetText := CreateTargetText()
	return model{
		userTotalInput: "",
		targetText:     targetText,
		cursor:         0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			char := msg.Runes[0]
			m.inputChar = char
			m.userTotalInput += string(char)

		case tea.KeySpace:
			m.inputChar = ' '
			m.userTotalInput += " "

		case tea.KeyBackspace:
			if len(m.userTotalInput) > 0 {
				m.userTotalInput = m.userTotalInput[:len(m.userTotalInput)-1]
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	}
	return m, nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ALAS! AN ERROR AHOY: %v", err)
		os.Exit(1)
	}
}
