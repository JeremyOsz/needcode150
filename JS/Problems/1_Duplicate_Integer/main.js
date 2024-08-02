class Solution {
    constructor() {
        // Initialize any properties if needed
    }

  /**
   * @param {number[]} nums
   * @return {boolean}
   */
  hasDuplicate(nums) {
      const seen = new Set();

      for(let i = 0; i < nums.length; i++){
          if (seen.has(nums[i])){
              console.log(true)
              return true
          }
          seen.add(nums[i]);
      }
      console.log(false)
      return false
  }
}

module.exports = Solution;
