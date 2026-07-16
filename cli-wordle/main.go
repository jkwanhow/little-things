package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
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

/*
func main() {
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
*/

type model struct {
	guesses        [6][5]string
	letterStates   [6][5]int
	curRow, curCol int
	answer         string
	dictionary     map[string]bool
	invalidState   string
}

const (
	Reset = "\033[0m"
	// The numbers define the color:
	// 42 = green background, 37 = white text
	Green  = "\033[42;37m"
	Yellow = "\033[43;37m"

	// 100 = dark gray background
	Gray = "\033[100;37m"
)

// View implements tea.Model.
func (m model) View() string {
	// The header
	s := "WORDLE?!\n\n"
	// Iterate over our choices
	for row, word := range m.guesses {
		rowState := m.letterStates[row]
		for col, letter := range word {
			colorState := Gray
			if rowState[col] == 1 {
				colorState = Green
			} else if rowState[col] == -1 {
				colorState = Yellow
			}

			s += fmt.Sprintf("%s[ %s ]%s", colorState, letter, Reset)
		}
		s += fmt.Sprint("\n")
	}

	if m.invalidState == "length" {
		s += "\n*Not enough letters*\n"
	} else if m.invalidState == "non-word" {
		s += "\n*Not in word list*\n"
	}

	// The footer
	s += "\nPress ctrl+c to quit.\n"

	// Send the UI for rendering
	return s
}

func initialModel() model {
	dictionary := CreateDictionary()
	return model{
		curRow:     0,
		curCol:     0,
		dictionary: dictionary,
		answer:     "boats",
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
			m.invalidState = ""
			if m.curCol > 0 {
				m.curCol--
				m.guesses[m.curRow][m.curCol] = ""
			}

		case "enter":
			splitGuess := m.guesses[m.curRow]
			var guess string
			for _, char := range splitGuess {
				guess += char
			}
			guess = CleanString(guess)
			if len(guess) != 5 {
				m.invalidState = "length"
			} else if !m.dictionary[guess] {
				m.invalidState = "non-word"
			} else {
				// need to check if the word is correct too.
				m.letterStates[m.curRow] = GetStatesOfLetters(m.answer, guess)
				// process the colors and all
				m.curRow++
				m.curCol = 0
			}

		default:
			m.invalidState = ""
			if m.curCol < 5 {
				keyStr := msg.String()
				if len(keyStr) == 1 && unicode.IsLetter(rune(keyStr[0])) {
					letter := keyStr
					m.guesses[m.curRow][m.curCol] = CleanString(letter)
					m.curCol++
				}
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
