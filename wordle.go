package main

import (
	"math/rand/v2"
	"slices"
	"strings"
	"unicode"

	// https://github.com/charmbracelet/lipgloss
	"charm.land/lipgloss/v2"
)

func isValidInput(userInput string) bool {
	if len(userInput) < 5 {
		return false
	}

	for _, r := range userInput {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func chooseWord(fetchNyt bool) string {
	if fetchNyt {
		return fetchNYTWord()
	}
	randomIndex := rand.IntN(len(words))
	return words[randomIndex]
}

func checkWord(answer string, testWord string) (string, bool) {
	answerArr := strings.Split(strings.ToUpper(answer), "")
	testWordArr := strings.Split(strings.ToUpper(testWord), "")

	correctStyle := lipgloss.NewStyle().Background(lipgloss.Color("#618C55")).Bold(true)
	incorrectStyle := lipgloss.NewStyle().Background(lipgloss.Color("#3A3A3C")).Bold(true)
	wrongSpotStyle := lipgloss.NewStyle().Background(lipgloss.Color("#B2A04C")).Bold(true)

	resultStr := make([]string, 5)

	for i, c := range testWordArr {
		resultStr[i] = incorrectStyle.Render(" ", c, " ")
	}

	correctWord := true
	// first iteration for correct
	for i, c := range testWordArr {
		if answerArr[i] == c {
			resultStr[i] = correctStyle.Render(" ", c, " ")
			answerArr[i] = "0"
		} else {
			correctWord = false
		}
	}

	// Second iteration for wrong spot
	for i, c := range testWordArr {
		if answerArr[i] == "0" {
			continue
		}
		if slices.Contains(answerArr, c) {
			resultStr[i] = wrongSpotStyle.Render(" ", c, " ")
			idx := slices.Index(answerArr, c)
			answerArr[idx] = "1"
		}
	}

	return strings.Join(resultStr, ""), correctWord
}
