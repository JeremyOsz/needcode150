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
	expected := "Hello, Two_Integer_Sum!\n"
	if buf.String() != expected {
		t.Errorf("expected %q but got %q", expected, buf.String())
	}
}

// Case 1
// Input:
// nums = [3,4,5,6], target = 7

// Output: [0,1]

// Case 2
// Input: nums = [4,5,6], target = 10
// Output: [0,2]

// Case 3
// Input: nums = [5,5], target = 10
// Output: [0,1]

// Intialize the test cases
var testCases = []struct {
	nums     []int
	target   int
	expected []int
}{
	{[]int{3, 4, 5, 6}, 7, []int{0, 1}},
	{[]int{4, 5, 6}, 10, []int{0, 2}},
	{[]int{5, 5}, 10, []int{0, 1}},
}

// TestTwoSum tests the TwoSum function against the test cases
func TestTwoSum(t *testing.T) {
	for _, tc := range testCases {
		result := TwoSum(tc.nums, tc.target)
		if !testutils.Equal(result, tc.expected) {
			t.Errorf("expected %v but got %v", tc.expected, result)
		}
	}
}
