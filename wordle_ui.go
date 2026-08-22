package main

import (
	"slices"
	"strings"

	// https://github.com/charmbracelet/bubbletea
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	title        string
	guesses      []string
	currentGuess int
}

func InitialModel() Model {
	defaultGuess := lipgloss.NewStyle().Background(lipgloss.Color("#3A3A3C")).Render(strings.Repeat("  .  ", 5))
	return Model{
		title:        "WORDLE",
		guesses:      slices.Repeat([]string{defaultGuess}, 6),
		currentGuess: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

var (
	answer = "PRAWN"
	words  = []string{
		"cream",
		"south",
		"craft",
		"blips",
		"grasp",
		"prawn",
	}
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.guesses[m.currentGuess] = validateWord(answer, words[m.currentGuess])
			m.currentGuess++
			if m.currentGuess > 5 {
				m.currentGuess = 5
			}
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	titleStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#618C55")).
		Padding(0, 10).
		Margin(1, 1).
		Align(lipgloss.Center).Bold(true)

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#606064")).
		Margin(0, 2)

	guessBlockStyle := lipgloss.NewStyle().
		Margin(1, 2)

	s := titleStyle.Render(m.title) + "\n"

	guesses := ""
	for _, word := range m.guesses {
		guesses += word + "\n"
	}

	s += guessBlockStyle.Render(guesses) + "\n"

	s += footerStyle.Render("Enter to Submit\nCtrl+c to Quit")

	s = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		Padding(1, 2).
		Render(s)

	return tea.NewView(s)
}
