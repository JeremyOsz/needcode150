#!/bin/bash

# A script to make a new problem
# Accepts arguments <problem_name> <problem_number> <language>

# Check if the user has provided the correct number of arguments
if [ $# -ne 3 ]; then
    echo "Usage: $0 <problem_name> <problem_number> <language>"
    exit 1
fi

# Check if the problem already exists
if [ -d "../$3/Problems/$2_$1" ]; then
    echo "Problem already exists"
    exit 1
fi

# Create the problem directory in the correct location
# Eg. For newProblem Duplicate_Integer 1 go to Go/Problems/1
filename=$2_$1
mkdir -p "../$3/Problems/$filename"

# Debugging: Check current directory
echo "Current directory: $(pwd)"

# For Go, create the main.go and main_test.go files
if [ "$3" = "Go" ]; then
    # Read the templates
    mainTemplate=$(cat "../Go/templates/main.go")
    testTemplate=$(cat "../Go/templates/main.test.go")

    # Replace placeholders with actual values
    mainContent=$(echo "$mainTemplate" | sed "s/{{.ProblemName}}/$1/")
    testContent=$(echo "$testTemplate" | sed "s/{{.ProblemName}}/$1/")

    # Write the main.go file
    echo "$mainContent" > "../$3/Problems/$filename/main.go"

    # Write the main_test.go file
    echo "$testContent" > "../$3/Problems/$filename/main_test.go"
fi

# For JavaScript, create the main.js and main.test.js files
if [ "$3" = "JS" ]; then
    # Read the templates
    mainTemplate=$(cat "../JS/templates/main.js")
    testTemplate=$(cat "../JS/templates/main.test.js")

    # Replace placeholders with actual values
    mainContent=$(echo "$mainTemplate" | sed "s/{{.ProblemName}}/$1/")
    testContent=$(echo "$testTemplate" | sed "s/{{.ProblemName}}/$1/")

    # Write the main.js file
    echo "$mainContent" > "../$3/Problems/$filename/main.js"

    # Write the main_test.js file
    echo "$testContent" > "../$3/Problems/$filename/main_test.js"
fi

# Debugging: Check if files are created
if [ -f "../$3/Problems/$filename/main.go" ] && [ -f "../$3/Problems/$filename/main_test.go" ]; then
    echo "Go files created successfully"
elif [ -f "../$3/Problems/$filename/main.js" ] && [ -f "../$3/Problems/$filename/main_test.js" ]; then
    echo "JavaScript files created successfully"
else
    echo "Failed to create files"
fi

echo "Problem $filename created successfully in $3/Problems/$filename"