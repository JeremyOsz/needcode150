class Solution {
  constructor() {
    // Initialize any properties if needed
  }

  helloWorld() {
    return "Hello, World!";
  }

  /**
   * @param {number[]} nums
   * @param {number} target
   * @return {number[]}
   */
  twoSum(nums, target) {
    // Interate through nums array to create a map
    const map = new Map();
    for (let i = 0; i < nums.length; i++) {
      const complement = target - nums[i];

      // Check if the complement is already in the map
      if (map.has(complement)) {
        return [map.get(complement), i];
      }

      // If the complement is not found, add the current number to the map
      map.set(nums[i], i);
    }
  }

  // Inefficient O(n^2) solution loop through the array twice
  twoSumInefficient(nums, target) {
    // Iterate through nums array
    for (let i = 0; i < nums.length; i++) {
      for (let j = i + 1; j < nums.length; j++) {
        // Check if the sum of nums[i] and nums[j] is equal to target
        if (nums[i] + nums[j] === target) {
          return [i, j];
        }
      }
    }
  }
}

module.exports = Solution;
