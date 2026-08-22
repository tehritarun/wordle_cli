package main

import (
	"slices"
	"strings"

	// https://github.com/charmbracelet/bubbletea
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	TITLESTYLE      = lipgloss.NewStyle().Background(lipgloss.Color("#618C55")).Padding(0, 10).Margin(1, 1).Align(lipgloss.Center).Bold(true)
	FOOTERSTYLE     = lipgloss.NewStyle().Foreground(lipgloss.Color("#707074")).Margin(0, 2)
	GUESSBLOCKSTYLE = lipgloss.NewStyle().Margin(1, 2)
	DEFAULTSTYLE    = lipgloss.NewStyle().Background(lipgloss.Color("#3A3A3C"))
	MESSAGESTYLE    = lipgloss.NewStyle().Foreground(lipgloss.Color("#618C55")).Align(lipgloss.Center).Margin(0, 2)
	FRAMESTYLE      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Padding(1, 2).Margin(1)
)

type Model struct {
	title         string
	guesses       []string
	currentGuess  int
	inputText     textinput.Model
	matchover     bool
	messageString string
	quitting      bool
}

func InitialModel() Model {
	defaultGuess := DEFAULTSTYLE.Render(strings.Repeat("  .  ", 5))

	ti := textinput.New()
	ti.Placeholder = "Enter word"
	ti.Focus()
	ti.CharLimit = 5
	ti.SetWidth(25)

	return Model{
		title:         "WORDLE",
		guesses:       slices.Repeat([]string{defaultGuess}, 6),
		currentGuess:  0,
		inputText:     ti,
		messageString: "",
		matchover:     false,
		quitting:      false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

var answer = "PRAWN"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q":
			if m.matchover {
				m.quitting = true
				return m, tea.Quit
			}

		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			if m.matchover {
				return m, nil
			}

			userInput := m.inputText.Value()
			if !isValidInput(userInput) {
				m.messageString = "Invalid Input"
				return m, nil
			}
			m.messageString = ""

			m.inputText.SetValue("")
			guess, matchWon := checkWord(answer, userInput)
			m.guesses[m.currentGuess] = guess

			if matchWon {
				m.matchover = true
				m.messageString = "You Won"
				return m, nil
			}

			if m.currentGuess >= 5 {
				m.currentGuess = 5
				m.matchover = true
				m.messageString = "You lost, Answer: " + answer
			}
			m.currentGuess++
		}
	}
	m.inputText, cmd = m.inputText.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	if m.quitting {
		return tea.NewView(MESSAGESTYLE.Render("Have a Nice Day!"))
	}

	s := TITLESTYLE.Render(m.title) + "\n"

	guesses := strings.Join(m.guesses, "\n")

	s += GUESSBLOCKSTYLE.Render(guesses) + "\n\n"

	if !m.matchover {
		s += m.inputText.View() + "\n\n"
	}

	s += MESSAGESTYLE.Render(m.messageString) + "\n\n"

	footer := "Press Ctrl+c to Quit\n"
	if m.matchover {
		footer += "Press Q to quit"
	} else {
		footer += "Press enter to Submit"
	}
	s += FOOTERSTYLE.Render(footer)

	s = FRAMESTYLE.Render(s)

	return tea.NewView(s)
}
