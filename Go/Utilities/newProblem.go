package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const goMainTemplate = `package main

import "fmt"

func main() {
    fmt.Println("Hello, {{.ProblemName}}!")
}
`

const goTestTemplate = `package main

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
    expected := "Hello, {{.ProblemName}}!\\n"
    if buf.String() != expected {
        t.Errorf("expected %q but got %q", expected, buf.String())
    }
}
`

type Problem struct {
	ProblemName   string
	ProblemNumber string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: <program> <problem_name> <problem_number>")
		os.Exit(1)
	}

	problem := Problem{
		ProblemName:   os.Args[1],
		ProblemNumber: os.Args[2],
	}

	dirPath := filepath.Join("../", fmt.Sprintf("%s_%s", problem.ProblemNumber, problem.ProblemName))
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		fmt.Println("Problem already exists")
		os.Exit(1)
	}

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Printf("Failed to create directory: %v\n", err)
		os.Exit(1)
	}

	createFileFromTemplate(filepath.Join(dirPath, "main.go"), goMainTemplate, problem)
	createFileFromTemplate(filepath.Join(dirPath, "main_test.go"), goTestTemplate, problem)

	fmt.Printf("Problem %s_%s created successfully in %s_%s\n", problem.ProblemNumber, problem.ProblemName, problem.ProblemNumber, problem.ProblemName)
}

func createFileFromTemplate(filePath, tmpl string, problem Problem) {
	t := template.Must(template.New("file").Parse(tmpl))
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	if err := t.Execute(file, problem); err != nil {
		fmt.Printf("Failed to write to file %s: %v\n", filePath, err)
		os.Exit(1)
	}
}
