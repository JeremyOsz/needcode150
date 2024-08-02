class Solution {
  constructor() {
    // Initialize any properties if needed
  }

  /**
   * @param {string} s
   * @param {string} t
   * @return {boolean}
   */
  isAnagram(s, t) {
    const sArr = s.split("").sort().join("");
    const tArr = t.split("").sort().join("");
    return sArr === tArr;
  }

  // For long strings or large inputs, a Frequency Counter is more efficient
  /**
   * @param {string} s
   * @param {string} t
   * @return {boolean}
   */
  isAnagramByFrequency(s, t) {
    if (s.length !== t.length) {
      return false;
    }

    const count = {};

    for (let char of s) {
      // Add 1 to the count if it exists, otherwise set it to 1
      count[char] = (count[char] || 0) + 1;
    }

    for (let char of t) {
      
      // If single character is not in second string, it cannot be an anagram
      if (!count[char]) {
        return false;
      }
      count[char]--;
    }

    return true;
  }
}

module.exports = Solution;
