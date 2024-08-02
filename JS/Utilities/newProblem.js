const fs = require('fs');
const path = require('path');
const { exec } = require('child_process');

const mainTemplate = `
module.exports = "Hello, Is_Anagram!";
`;

const testTemplate = `
const Solution = require('./main'); // Adjust the path if necessary

test('main function output', () => {
    const output = Solution
    console.log(output)
    const expected = "Hello, Is_Anagram!";
    expect(output).toBe(expected);
});
`;

function createFileFromTemplate(filePath, template, problem) {
    const content = template.replace(/{{ProblemName}}/g, problem.problemName);
    fs.writeFileSync(filePath, content, 'utf8');
}

function main() {
    const args = process.argv.slice(2);
    if (args.length !== 2) {
        console.log('Usage: <program> <problem_name> <problem_number>');
        process.exit(1);
    }

    const problem = {
        problemName: args[0],
        problemNumber: args[1]
    };

    const dirPath = path.join(__dirname, '..', 'Problems', `${problem.problemNumber}_${problem.problemName}`);
    if (fs.existsSync(dirPath)) {
        console.log('Problem already exists');
        process.exit(1);
    }

    fs.mkdirSync(dirPath, { recursive: true });

    createFileFromTemplate(path.join(dirPath, 'main.js'), mainTemplate, problem);
    createFileFromTemplate(path.join(dirPath, 'main.test.js'), testTemplate, problem);

    console.log(`Problem ${problem.problemNumber}_${problem.problemName} created successfully in ${problem.problemNumber}_${problem.problemName}`);
}

main();