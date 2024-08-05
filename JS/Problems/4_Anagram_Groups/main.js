class Solution {
  constructor() {
    // Initialize any properties if needed
  }

  helloWorld() {
    return "Hello, World!";
  }

  groupAnagramsSort(strs) {
    const map = new Map();
    for (let str of strs) {
      const sortedStr = str.split("").sort().join("");

      // Check if the key exists in the map
      if (!map.has(sortedStr)) {
        // Initialize the array if the key does not exist
        map.set(sortedStr, []);
      }

      // Push the value to the array
      map.get(sortedStr).push(str);
    }
    // make map values into an array
    const result = Array.from(map.values());
    return result;
  }

  /**
     * @param {string[]} strs
     * @return {string[][]}
     */
  groupAnagrams(strs) {
    const map = new Map();

    for (let str of strs) {
        // Create a frequency count of characters
        const charCount = new Array(26).fill(0);
        for (let char of str) {
            charCount[char.charCodeAt(0) - 'a'.charCodeAt(0)]++;
        }

        // Convert the frequency count to a string key
        const key = charCount.join('#');
        console.log(key)

        // Check if the key exists in the map
        if (!map.has(key)) {
            // Initialize the array if the key does not exist
            map.set(key, []);
        }

        // Push the value to the array
        map.get(key).push(str);

    }

    // Make map values into an array
    const result = Array.from(map.values());
    return result;
}
}

module.exports = Solution;
