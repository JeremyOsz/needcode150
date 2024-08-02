const Solution = require('./main'); // Adjust the path if necessary

// import 

test('main function output', () => {
    const output = Solution
    console.log(output)
    const expected = "Hello, Is_Anagram!";
    expect(output).toBe(expected);
});
