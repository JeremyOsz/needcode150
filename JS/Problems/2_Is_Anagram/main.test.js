const Solution = require('./main'); // Adjust the path if necessary

// import 

test('main function output', () => {
    const output = Solution
    console.log(output)
    const expected = "Hello, Is_Anagram!";
    expect(output).toBe(expected);
});

// Eg. 1
// Input: s = racecar, t = carrace 
// Output: true
test('anagram should return true', () => {
    const solution = new Solution();
    const output = solution.isAnagram("racecar", "carrace")
    console.log(output)
    const expected = true;
    expect(output).toBe(expected);
});

// Eg. 2
// Input: s = "jar", t = "jam"
// Output: false
test('anagram should return false', () => {
    const solution = new Solution();
    const output = solution.isAnagram("jar", "jam")
    console.log(output)
    const expected = false;
    expect(output).toBe(expected);
});