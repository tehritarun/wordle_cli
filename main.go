package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(validateWord("RETRO", "ETROR"))
	fmt.Println(validateWord("RETRO", "MATRO"))
	fmt.Println(validateWord("RETRO", "RTROE"))
	fmt.Println(validateWord("RETRO", "RETRE"))
}

func validateWord(answer string, testWord string) string {
	answerArr := strings.Split(strings.ToUpper(answer), "")
	testWordArr := strings.Split(strings.ToUpper(testWord), "")

	correct := "🟩"
	incorrect := "⬛️"
	wrongSpot := "🟨"

	result := slices.Repeat([]string{incorrect}, 5)

	// first iteration for correct
	for i, c := range testWordArr {
		if answerArr[i] == c {
			result[i] = correct
			answerArr[i] = "0"
		}
	}

	for i, c := range testWordArr {
		if answerArr[i] == "0" {
			continue
		}
		if slices.Contains(answerArr, c) {
			result[i] = wrongSpot
			idx := slices.Index(answerArr, c)
			answerArr[idx] = "1"
		}
	}

	return strings.Join(result, "")
}
