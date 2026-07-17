package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

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

type model struct {
	guesses        [6][5]string
	letterStates   [6][5]int
	curRow, curCol int
	answer         string
	dictionary     map[string]bool
	message        string
	state          string
	fontMap        map[string][]string
}

const (
	Reset  = "\033[0m"
	Green  = "\033[42;37m"
	Yellow = "\033[43;37m"
	Orange = "\033[48;5;208;37m"
	Gray   = "\033[100;37m"
)

// View implements tea.Model.
func (m model) View() string {
	// The header
	s := "WORDLE?!\n\n"
	s += ""
	s += "\n\n"

	for row, word := range m.guesses {
		rowState := m.letterStates[row]
		for i := 0; i < HEIGHT; i++ {
			for col, letter := range word {
				fontLetterLines, ok := m.fontMap[strings.ToUpper(letter)]
				var fontLine string
				if ok {
					fontLine = fontLetterLines[i]
				} else {
					fontLine = EMPTYLINE
				}
				colorState := Gray
				switch rowState[col] {
				case 1:
					colorState = Green
				case -1:
					colorState = Orange
				default:
					colorState = Gray
				}

				s += fmt.Sprintf("%s%s%s%s%s", SPACE, colorState, fontLine, Reset, SPACE)
			}
			s += "\n"
		}

		s += "\n"
	}

	if m.state != "play" {
		s += fmt.Sprintf("\n*%s*\n", m.message)
	}

	// The footer
	s += "\nPress ctrl+c to quit.\n"

	// Send the UI for rendering
	return s
}

func initialModel() model {
	dictionary := CreateDictionary()
	fontMap := CreateFontMap()
	return model{
		curRow:     0,
		curCol:     0,
		dictionary: dictionary,
		answer:     "boats",
		state:      "play",
		fontMap:    fontMap,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	//key is pressed
	case tea.KeyMsg:

		// what key is pressed
		switch msg.String() {

		// close or exit the program
		case "ctrl+c":
			return m, tea.Quit

		case "backspace":
			if m.state != "end" {
				m.HandleDelete()
			}

		case "enter":
			if m.state != "end" {
				m.EnterLine()
			}

		default:
			if m.state != "end" {
				m.CharacterInput(msg)
			}
		}
	}

	return m, nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
