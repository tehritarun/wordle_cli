package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	// validateWord("RETRO", "ETROR")
	// validateWord("RETRO", "MATRO")
	// validateWord("RETRO", "RTROE")
	// validateWord("RETRO", "RETRE")
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
