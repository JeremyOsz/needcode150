const Solution = require('./main'); // Adjust the path if necessary

test('main function output', () => {
    const solution = new Solution();
    const output = solution.helloWorld()
    console.log(output)
    const expected = "Hello, World!";
    expect(output).toBe(expected);
});