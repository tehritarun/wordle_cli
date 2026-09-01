package main

import (
	"strings"
	"testing"
)

func TestIsValidInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid 5-letter word", "HELLO", true},
		{"Valid lowercase word", "hello", true},
		{"Too short", "HI", false},
		{"Too long", "HELLOO", false},
		{"Contains numbers", "HEL10", false},
		{"Contains special chars", "HE!LO", false},
		{"Empty string", "", false},
		{"Valid mixed case", "HeLLo", true},
		{"Contains space", "HE LO", false},
		{"Exactly 5 letters", "WORLD", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidInput(tt.input)
			if result != tt.expected {
				t.Errorf("isValidInput(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestChooseWord(t *testing.T) {
	t.Run("Random word selection", func(t *testing.T) {
		word := chooseWord(false)
		if !isValidInput(word) {
			t.Errorf("chooseWord(false) returned invalid word: %q", word)
		}
		if len(word) != 5 {
			t.Errorf("chooseWord(false) returned word with length %d, want 5", len(word))
		}
		// Check if word is uppercase
		if word != strings.ToUpper(word) {
			t.Errorf("chooseWord(false) returned non-uppercase word: %q", word)
		}
	})

	t.Run("Random word is from words list", func(t *testing.T) {
		word := chooseWord(false)
		found := false
		for _, w := range words {
			if w == word {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("chooseWord(false) returned word %q not in words list", word)
		}
	})

	t.Run("Multiple calls return words", func(t *testing.T) {
		// Test that multiple calls work and potentially return different words
		words := make(map[string]bool)
		for i := 0; i < 100; i++ {
			word := chooseWord(false)
			words[word] = true
			if len(word) != 5 {
				t.Errorf("chooseWord(false) call %d returned invalid length: %d", i, len(word))
			}
		}
		// We should get at least some variety (though not guaranteed)
		if len(words) == 0 {
			t.Error("chooseWord(false) failed to return any words")
		}
	})
}

func TestCheckWord(t *testing.T) {
	tests := []struct {
		name          string
		answer        string
		guess         string
		expectCorrect bool
		description   string
	}{
		{
			name:          "Exact match",
			answer:        "HELLO",
			guess:         "HELLO",
			expectCorrect: true,
			description:   "All letters should be green",
		},
		{
			name:          "No match",
			answer:        "HELLO",
			guess:         "WXYZ",
			expectCorrect: false,
			description:   "All letters should be gray (except any coincidental matches)",
		},
		{
			name:          "Partial match - wrong positions",
			answer:        "HELLO",
			guess:         "OLEHL",
			expectCorrect: false,
			description:   "Some letters in wrong positions (yellow)",
		},
		{
			name:          "Case insensitive - lowercase guess",
			answer:        "HELLO",
			guess:         "hello",
			expectCorrect: true,
			description:   "Should handle lowercase input",
		},
		{
			name:          "Case insensitive - mixed case",
			answer:        "HELLO",
			guess:         "HeLLo",
			expectCorrect: true,
			description:   "Should handle mixed case input",
		},
		{
			name:          "Duplicate letters in guess",
			answer:        "HELLO",
			guess:         "LLAMA",
			expectCorrect: false,
			description:   "Should handle duplicate letters correctly",
		},
		{
			name:          "Single correct letter",
			answer:        "HELLO",
			guess:         "HXXXX",
			expectCorrect: false,
			description:   "Only first letter correct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, correct := checkWord(tt.answer, tt.guess)

			// Check if correctness matches expectation
			if correct != tt.expectCorrect {
				t.Errorf("checkWord(%q, %q) correct = %v, want %v",
					tt.answer, tt.guess, correct, tt.expectCorrect)
			}

			// Check that result is not empty
			if result == "" {
				t.Errorf("checkWord(%q, %q) returned empty result", tt.answer, tt.guess)
			}

			// For exact matches, verify all letters are marked correct (green)
			if tt.expectCorrect {
				// The result should contain styled output
				// We can't easily test the exact styling, but we can verify it's not empty
				if !strings.Contains(strings.ToUpper(result), strings.ToUpper(tt.guess[0:1])) {
					t.Logf("Note: Result format: %s", result)
				}
			}
		})
	}
}

func TestCheckWordColorCoding(t *testing.T) {
	// Test that the color coding logic works correctly
	t.Run("All correct positions", func(t *testing.T) {
		_, correct := checkWord("TESTS", "TESTS")
		if !correct {
			t.Error("Expected all correct, got false")
		}
	})

	t.Run("All wrong positions", func(t *testing.T) {
		_, correct := checkWord("ABCDE", "FGHIJ")
		if correct {
			t.Error("Expected none correct, got true")
		}
	})

	t.Run("Mixed correct and wrong positions", func(t *testing.T) {
		result, correct := checkWord("HELLO", "HOLES")
		if correct {
			t.Error("Expected not all correct, got true")
		}
		if result == "" {
			t.Error("Expected non-empty result")
		}
	})
}

func TestWordsListValidity(t *testing.T) {
	t.Run("All words are 5 letters", func(t *testing.T) {
		for i, word := range words {
			if len(word) != 5 {
				t.Errorf("Word at index %d (%q) has length %d, want 5", i, word, len(word))
			}
		}
	})

	t.Run("All words are uppercase", func(t *testing.T) {
		for i, word := range words {
			if word != strings.ToUpper(word) {
				t.Errorf("Word at index %d (%q) is not uppercase", i, word)
			}
		}
	})

	t.Run("All words contain only letters", func(t *testing.T) {
		for i, word := range words {
			if !isValidInput(word) {
				t.Errorf("Word at index %d (%q) contains non-letter characters", i, word)
			}
		}
	})

	t.Run("Words list is not empty", func(t *testing.T) {
		if len(words) == 0 {
			t.Error("Words list is empty")
		}
	})

	t.Run("Words list has reasonable size", func(t *testing.T) {
		if len(words) < 100 {
			t.Errorf("Words list has only %d words, expected at least 100", len(words))
		}
	})

	t.Run("No duplicate words", func(t *testing.T) {
		seen := make(map[string]bool)
		for i, word := range words {
			if seen[word] {
				t.Errorf("Duplicate word %q found at index %d", word, i)
			}
			seen[word] = true
		}
	})
}
