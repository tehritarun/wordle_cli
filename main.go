package main

import (
	"fmt"
	"slices"
	"strings"

	"charm.land/lipgloss/v2"
)

func main() {
	validateWord("RETRO", "ETROR")
	validateWord("RETRO", "MATRO")
	validateWord("RETRO", "RTROE")
	validateWord("RETRO", "RETRE")
}

func validateWord(answer string, testWord string) string {
	answerArr := strings.Split(strings.ToUpper(answer), "")
	testWordArr := strings.Split(strings.ToUpper(testWord), "")

	correct := "🟩"
	incorrect := "⬛️"
	wrongSpot := "🟨"

	correctStyle := lipgloss.NewStyle().Background(lipgloss.Color("#6CA965"))
	incorrectStyle := lipgloss.NewStyle().Background(lipgloss.Color("#787C7F"))
	wrongSpotStyle := lipgloss.NewStyle().Background(lipgloss.Color("#C8B653"))

	result := slices.Repeat([]string{incorrect}, 5)
	resultStr := make([]string, 5)

	for i, c := range testWordArr {
		resultStr[i] = incorrectStyle.Render(" ", c, " ")
	}

	// first iteration for correct
	for i, c := range testWordArr {
		if answerArr[i] == c {
			result[i] = correct
			resultStr[i] = correctStyle.Render(" ", c, " ")
			answerArr[i] = "0"
		}
	}

	// Second iteration for wrong spot
	for i, c := range testWordArr {
		if answerArr[i] == "0" {
			continue
		}
		if slices.Contains(answerArr, c) {
			result[i] = wrongSpot
			resultStr[i] = wrongSpotStyle.Render(" ", c, " ")
			idx := slices.Index(answerArr, c)
			answerArr[idx] = "1"
		}
	}

	fmt.Println(strings.Join(resultStr, ""))
	return strings.Join(result, "")
}
