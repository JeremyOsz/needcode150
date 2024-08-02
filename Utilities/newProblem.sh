#!/bin/bash

# A script to make a new problem
# Accepts arguments <problem_name> <problem_number> <language>

# Check if the user has provided the correct number of arguments
if [ $# -ne 3 ]; then
    echo "Usage: $0 <problem_name> <problem_number> <language>"
    exit 1
fi

# Check if the problem already exists
if [ -d ../$3/Problems/$2_$1 ]; then
    echo "Problem already exists"
    exit 1
fi

# Create the problem directory in the correct location
# Eg. For newProblem Duplicate_Integer 1 go to Go/Problems/1
filename=$2_$1
mkdir -p ../$3/Problems/$filename
cd ../$3/Problems/$filename

# Debugging: Check current directory
echo "Current directory: $(pwd)"

# For Go, create the main.go and main_test.go files
if [ "$3" = "Go" ]; then
    cat <<EOL > main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, $1!")
}
EOL

    cat <<EOL > main_test.go
package main

import (
    "bytes"
    "io"
    "os"
    "testing"
)

func TestMain(t *testing.T) {
    // Save the original stdout
    originalStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Call the main function
    main()

    // Capture the output
    w.Close()
    var buf bytes.Buffer
    io.Copy(&buf, r)
    os.Stdout = originalStdout

    // Check the output
    expected := "Hello, $1!\\n"
    if buf.String() != expected {
        t.Errorf("expected %q but got %q", expected, buf.String())
    }
}
EOL
fi


# For JavaScript, create the main.js and main.test.js files
if [ "$3" = "JS" ]; then
    cat <<EOL > main.js
module.exports = "Hello, Is_Anagram!";
EOL

    cat <<EOL > main.test.js
const Solution = require('./main'); // Adjust the path if necessary

test('main function output', () => {
    const output = Solution
    console.log(output)
    const expected = "Hello, Is_Anagram!";
    expect(output).toBe(expected);
});

EOL
fi

# Debugging: Check if files are created
if [ -f main.go ] && [ -f main_test.go ]; then
    echo "Go files created successfully"
elif [ -f main.js ] && [ -f main.test.js ]; then
    echo "JavaScript files created successfully"
else
    echo "Failed to create files"
fi

echo "Problem $filename created successfully in $3/Problems/$filename"