const Solution = require('./main'); // Adjust the path if necessary

test('main function output', () => {
    const solution = new Solution();
    const output = solution.helloWorld()
    console.log(output)
    const expected = "Hello, World!";
    expect(output).toBe(expected);
});

// eg. 1
// Input: strs = ["act","pots","tops","cat","stop","hat"]

// Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
test('Eg. 1', () => {
    const solution = new Solution();
    const output = solution.groupAnagrams(["act","pots","tops","cat","stop","hat"])
    console.log(output)
    const expected = [["act","cat"],["pots","tops","stop"],["hat"]];
    expect(output).toStrictEqual(expected);
});

// eg. 2
// Input: strs = ["x"]
// Output: [["x"]]

test('Eg. 2', () => {
    const solution = new Solution();
    const output = solution.groupAnagrams(["x"])
    console.log(output)
    const expected = [["x"]];
    expect(output).toStrictEqual(expected);
})

// eg. 3
// Input: strs = [""]

// Output: [[""]]
test('Eg. 3', () => {
    const solution = new Solution();
    const output = solution.groupAnagrams([""])
    console.log(output)
    const expected = [[""]];
    expect(output).toStrictEqual(expected);
})
