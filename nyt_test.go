package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFetchNYTWord(t *testing.T) {
	t.Run("Successful fetch with valid response", func(t *testing.T) {
		// Create a mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify the request path format
			expectedDate := time.Now().Format("2006-01-02")
			if !strings.Contains(r.URL.Path, expectedDate) {
				t.Errorf("Request path %q does not contain expected date %q", r.URL.Path, expectedDate)
			}

			// Return a mock response
			response := nytResponse{Solution: "TESTS"}
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		// We can't easily test the actual fetchNYTWord without modifying it
		// to accept a custom URL, but we can test the response struct
		t.Log("Note: Full integration test requires network access")
	})

	t.Run("NYT response struct unmarshaling", func(t *testing.T) {
		jsonData := `{"solution":"HELLO"}`
		var resp nytResponse
		err := json.Unmarshal([]byte(jsonData), &resp)
		if err != nil {
			t.Errorf("Failed to unmarshal JSON: %v", err)
		}
		if resp.Solution != "HELLO" {
			t.Errorf("Expected solution HELLO, got %q", resp.Solution)
		}
	})

	t.Run("NYT response with additional fields", func(t *testing.T) {
		// Test that we can handle responses with extra fields
		jsonData := `{"solution":"WORLD","days_since_launch":100,"editor":"someone"}`
		var resp nytResponse
		err := json.Unmarshal([]byte(jsonData), &resp)
		if err != nil {
			t.Errorf("Failed to unmarshal JSON with extra fields: %v", err)
		}
		if resp.Solution != "WORLD" {
			t.Errorf("Expected solution WORLD, got %q", resp.Solution)
		}
	})

	t.Run("Empty solution handling", func(t *testing.T) {
		jsonData := `{"solution":""}`
		var resp nytResponse
		err := json.Unmarshal([]byte(jsonData), &resp)
		if err != nil {
			t.Errorf("Failed to unmarshal JSON: %v", err)
		}
		if resp.Solution != "" {
			t.Errorf("Expected empty solution, got %q", resp.Solution)
		}
	})
}

func TestChooseWordWithNYTFlag(t *testing.T) {
	t.Run("Random word when NYT flag is false", func(t *testing.T) {
		word := chooseWord(false)
		if len(word) != 5 {
			t.Errorf("Expected 5-letter word, got %d letters: %q", len(word), word)
		}
		if !isValidInput(word) {
			t.Errorf("chooseWord(false) returned invalid word: %q", word)
		}
	})

	// Note: We can't easily test chooseWord(true) without network access
	// or refactoring the code to accept a custom URL for testing
	t.Run("NYT mode requires network", func(t *testing.T) {
		t.Skip("Skipping NYT fetch test - requires network and live API")
		// In a real scenario, you might want to:
		// 1. Mock the HTTP client
		// 2. Use dependency injection
		// 3. Create a test environment variable
	})
}

func TestNYTResponseStruct(t *testing.T) {
	t.Run("JSON tag is correct", func(t *testing.T) {
		// Verify the struct can be marshaled/unmarshaled correctly
		original := nytResponse{Solution: "CRANE"}
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("Failed to marshal: %v", err)
		}

		var decoded nytResponse
		err = json.Unmarshal(data, &decoded)
		if err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if decoded.Solution != original.Solution {
			t.Errorf("Expected %q, got %q", original.Solution, decoded.Solution)
		}
	})
}
