package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func readTemplate(fileName string) string {
	content, err := os.ReadFile(filepath.Join("../templates", fileName))
	if err != nil {
		log.Fatalf("Failed to read template file: %v", err)
	}
	return string(content)
}

var goMainTemplate = readTemplate("main.go")

var goTestTemplate = readTemplate("main.test.go")

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

	dirPath := filepath.Join("../Problems/", fmt.Sprintf("%s_%s", problem.ProblemNumber, problem.ProblemName))
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
