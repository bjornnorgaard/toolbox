package oda

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReplaceAllCharsWith(filePath string, methodName string, r rune) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	method, err := findMethodWithName(methodName, string(file))

	parsedLines := replaceAllChars(method, r)

	return parsedLines, nil
}

func replaceAllChars(method []string, r rune) string {
	regex := regexp.MustCompile("\\S")
	for i, line := range method {
		method[i] = regex.ReplaceAllString(line, string(r))
	}

	parsedLines := strings.Join(method, "\n")
	return parsedLines
}

func findMethodWithName(methodName string, content string) ([]string, error) {
	var (
		lines           = strings.SplitN(content, "\n", -1)
		functionMatcher = fmt.Sprintf("func %s(", methodName)
		methodMatcher   = fmt.Sprintf(") func %s(", methodName)
		startIndex      = 0
		endIndex        = 0
		bracketCount    = 0
	)

	for i, line := range lines {
		if strings.Contains(line, functionMatcher) || strings.Contains(line, methodMatcher) {
			startIndex = i
			break
		}
	}

	for i := startIndex; i <= len(lines); i++ {
		bracketCount += strings.Count(lines[i], "{")
		bracketCount -= strings.Count(lines[i], "}")
		if bracketCount == 0 {
			endIndex = i
			break
		}
	}

	foundMethod := lines[startIndex : endIndex+1]
	return foundMethod, nil
}
