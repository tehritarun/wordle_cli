package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	filepath := os.Getenv("XDG_DATA_HOME")
	if filepath == "" {
		filepath = "~/.local/share"
	}
	filepath += "/wordle_cli/words.txt"
	chosenWord := chooseWord(filepath)
	if !isValidInput(chosenWord) {
		panic("Invalid word chosen from word file")
	}

	p := tea.NewProgram(InitialModel(chosenWord))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
