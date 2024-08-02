package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, Is_Anagram!")
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	// Make rune map for s
	for _, c := range s {
		m[c]++
	}

	// Decrement rune map for t

	for _, c := range t {
		if m[c] == 0 {
			return false
		}
		m[c]--
	}
	return true
}

// Solution using sort
func isAnagramSort(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := []rune(s)
	t1 := []rune(t)

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})

	for i := 0; i < len(s1); i++ {
		if s1[i] != t1[i] {
			return false
		}
	}
	return true
}
