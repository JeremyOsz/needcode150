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
	expected := "Hello, {{.ProblemName}}!\n"
	if buf.String() != expected {
		t.Errorf("expected %q but got %q", expected, buf.String())
	}
}
