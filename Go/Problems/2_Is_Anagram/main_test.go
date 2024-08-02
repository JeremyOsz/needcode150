package main

import (
	"bytes"
	"io"
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
	expected := "Hello, Is_Anagram!\n"
	if buf.String() != expected {
		t.Errorf("expected %q but got %q", expected, buf.String())
	}
}

// Eg. 1
// Input: s = racecar, t = carrace
// Output: true
func TestIsAnagram(t *testing.T) {
	a := "racecar"
	b := "carrace"
	expected := true
	got := isAnagram(a, b)
	if got != expected {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

// Test using sort method
func TestIsAnagramSort(t *testing.T) {
	a := "racecar"
	b := "carrace"
	expected := true
	got := isAnagramSort(a, b)
	if got != expected {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

// Eg. 2
// Input: s = "jar", t = "jam"
// Output: false
func TestIsNotAnagram(t *testing.T) {
	a := "jar"
	b := "jam"
	expected := false
	got := isAnagram(a, b)
	if got != expected {
		t.Errorf("expected %t but got %t", expected, got)
	}
}

func TestIsNotAnagramSort(t *testing.T) {
	a := "jar"
	b := "jam"
	expected := false
	got := isAnagramSort(a, b)
	if got != expected {
		t.Errorf("expected %t but got %t", expected, got)
	}
}
