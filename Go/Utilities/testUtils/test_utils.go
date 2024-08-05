// Package utilities provides utility functions for various operations.
package testutils

import "sort"

// Equal checks if two slices are equal.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// CompareStringSlices compares two slices of slices of strings, allowing any order
func CompareStringSlices(output [][]string, s [][]string) bool {
	if len(output) != len(s) {
		return false
	}

	// Sort the outer slices
	sort.Slice(output, func(i, j int) bool {
		return len(output[i]) < len(output[j])
	})
	sort.Slice(s, func(i, j int) bool {
		return len(s[i]) < len(s[j])
	})

	for i := 0; i < len(output); i++ {
		// Sort the inner slices
		sort.Strings(output[i])
		sort.Strings(s[i])

		if len(output[i]) != len(s[i]) {
			return false
		}

		for j := 0; j < len(output[i]); j++ {
			if output[i][j] != s[i][j] {
				return false
			}
		}
	}

	return true
}
