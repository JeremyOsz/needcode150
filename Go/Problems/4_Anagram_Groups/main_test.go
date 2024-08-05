package main

import (
	"bytes"
	"io"
	testutils "needcode-150/Utilities/testUtils"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Save the original stdout
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the main function
	main()

	// Capture the output
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = originalStdout

	// Check the output
	expected := "Hello, Anagram_Groups!\n"
	if buf.String() != expected {
		t.Errorf("expected %q but got %q", expected, buf.String())
	}
}

// eg. 1
// Input: strs = ["act","pots","tops","cat","stop","hat"]

// Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

// eg. 2
// Input: strs = ["x"]
// Output: [["x"]]

// eg. 3
// Input: strs = [""]

// Output: [[""]]

func TestGroupAnagram(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected [][]string
	}{
		{[]string{"act", "pots", "tops", "cat", "stop", "hat"}, [][]string{{"hat"}, {"pots", "tops", "stop"}, {"act", "cat"}}},
		{[]string{"x"}, [][]string{{"x"}}},
		{[]string{""}, [][]string{{""}}},
	}

	for _, tc := range testCases {
		output := GroupAnagram(tc.strs)
		if !testutils.CompareStringSlices(output, tc.expected) {
			t.Errorf("expected %v but got %v", tc.expected, output)
		}
	}
}
