package main

import (
	"slices"

	// https://github.com/charmbracelet/bubbletea
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	title        string
	guesses      []string
	currentGuess int
}

func InitialModel() Model {
	return Model{
		title:        "WORDLE",
		guesses:      slices.Repeat([]string{"               "}, 6),
		currentGuess: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	return tea.NewView("Hello, World!")
}
