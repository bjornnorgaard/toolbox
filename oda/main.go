package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	err := start("./main.go", "start")
	if err != nil {
		panic(err)
	}
}

func start(filePath string, methodName string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	method, err := findMethod(methodName, string(file))

	fmt.Println(method)
	return nil
}

func findMethod(method string, content string) (string, error) {
	lines := strings.SplitN(content, "\n", -1)

	functionMatcher := fmt.Sprintf("func %s(", method)
	methodMatcher := fmt.Sprintf(") func %s(", method)

	startIndex := 0
	endIndex := 0

	for i, line := range lines {
		if strings.Contains(line, functionMatcher) || strings.Contains(line, methodMatcher) {
			startIndex = i
			break
		}
	}

	bracketCount := 0
	for i := startIndex; i <= len(lines); i++ {
		bracketCount += strings.Count(lines[i], "{")
		bracketCount -= strings.Count(lines[i], "}")
		if bracketCount == 0 {
			endIndex = i + 1
			break
		}
	}

	regex := regexp.MustCompile("\\S")
	methodContent := lines[startIndex:endIndex]
	for i, line := range methodContent {
		methodContent[i] = regex.ReplaceAllString(line, "X")
	}

	parsedLines := strings.Join(methodContent, "\n")
	fmt.Printf("\n%s\n", parsedLines)

	return "", nil
}
