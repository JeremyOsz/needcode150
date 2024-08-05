package main

import "fmt"

func main() {
	fmt.Println("Hello, Two_Integer_Sum!")
}

func TwoSum(nums []int, target int) []int {
	// Initialize a map to store the indices of the numbers
	numMap := make(map[int]int)

	// Iterate through the slice
	for i, num := range nums {
		complement := target - num

		// Check if the complement is already in the map
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}

		// If the complement is not found, add the current number to the map
		numMap[num] = i
		fmt.Println(numMap)
	}

	// Return an empty slice if no solution is found
	return []int{}
}
