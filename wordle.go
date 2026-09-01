package main

import (
	"bufio"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"
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

func chooseWord(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic("enable to read file with words" + err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pickedWord string
	lineNum := 0
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Replace current pick with probability 1/lineNum
		if randGen.Intn(lineNum) == 0 {
			pickedWord = line
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return pickedWord
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
