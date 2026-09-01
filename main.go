package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	fetchNyt := false
	if len(os.Args) > 1 {
		fetchNyt = os.Args[1] == "-n"
	}
	chosenWord := chooseWord(fetchNyt)
	if !isValidInput(chosenWord) {
		panic("Invalid word chosen")
	}

	p := tea.NewProgram(InitialModel(chosenWord))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
