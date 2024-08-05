package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, Anagram_Groups!")
}

func GroupAnagram(strs []string) [][]string {
	// Create a map to store the grouped anagrams
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		// Create a frequency count of characters
		charCount := make([]int, 26)
		for _, char := range str {
			charCount[char-'a']++
		}

		fmt.Println(charCount)

		// Convert the frequency count to a string key
		var builder strings.Builder
		for _, count := range charCount {
			builder.WriteString(fmt.Sprintf("%d#", count))
		}
		key := builder.String()

		// Check if the key exists in the map
		if _, exists := anagramMap[key]; !exists {
			// Initialize the array if the key does not exist
			anagramMap[key] = []string{}
		}

		// Push the value to the array
		anagramMap[key] = append(anagramMap[key], str)
	}

	// Make map values into a slice of slices of strings
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}
