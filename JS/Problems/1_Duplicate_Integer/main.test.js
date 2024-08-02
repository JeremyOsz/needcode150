// test.js
const Solution = require('./main'); // Adjust the path if necessary

describe('Solution', () => {
    let solution;

    beforeEach(() => {
        solution = new Solution();
    });

    test('should return true for array [1, 2, 3, 3]', () => {
        expect(solution.hasDuplicate([1, 2, 3, 3])).toBe(true);
    });

    test('should return false for array [1, 2, 3, 4]', () => {
        expect(solution.hasDuplicate([1, 2, 3, 4])).toBe(false);
    });
});