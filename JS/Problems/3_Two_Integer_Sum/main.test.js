
const Solution = require('./main'); // Adjust the path if necessary

test('main function output', () => {
    const solution = new Solution();
    const output = solution.helloWorld()
    console.log(output)
    const expected = "Hello, World!";
    expect(output).toBe(expected);
});

// Constraints:

// 2 <= nums.length <= 1000
// -10,000,000 <= nums[i] <= 10,000,000
// -10,000,000 <= target <= 10,000,000

// Case 1
// Input: 
// nums = [3,4,5,6], target = 7

// Output: [0,1]
test('twoSum should return [0,1]', () => {
    const solution = new Solution();
    const output = solution.twoSum([3,4,5,6], 7)
    console.log(output)
    const expected = [0,1];
    expect(output).toStrictEqual(expected);
});

test('Inneficient twoSum should return [0,1]', () => {
    const solution = new Solution();
    const output = solution.twoSumInefficient([3,4,5,6], 7)
    console.log(output)
    const expected = [0,1];
    expect(output).toStrictEqual(expected);
});



// Case 2
// Input: nums = [4,5,6], target = 10
// Output: [0,2]
test('twoSum should return [0,2]', () => {
    const solution = new Solution();
    const output = solution.twoSum([4,5,6], 10)
    console.log(output)
    const expected = [0,2];
    expect(output).toStrictEqual(expected);
})

test('Inefficient twoSum should return [0,2]', () => {
    const solution = new Solution();
    const output = solution.twoSumInefficient([4,5,6], 10)
    console.log(output)
    const expected = [0,2];
    expect(output).toStrictEqual(expected);
})

// Output: [0,2]

// Case 3
// Input: nums = [5,5], target = 10
// Output: [0,1]
test('twoSum should return [0,1]', () => {
    const solution = new Solution();
    const output = solution.twoSum([5,5], 10)
    console.log(output)
    const expected = [0,1];
    expect(output).toStrictEqual(expected);
})

test('Inefficient twoSum should return [0,1]', () => {
    const solution = new Solution();
    const output = solution.twoSumInefficient([5,5], 10)
    console.log(output)
    const expected = [0,1];
    expect(output).toStrictEqual(expected);
})