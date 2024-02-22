package iniparse

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type INIResult map[string]Section

func Parse(file *os.File) (INIResult, error) {
	lines, err := readFile(file)
	if err != nil {
		return nil, err
	}
	sections := make(INIResult)

	currentSection := ""

	for idx, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.ReplaceAll(
				strings.ReplaceAll(line, "[", ""),
				"]",
				"",
			)
			sections[currentSection] = make(map[string]string)
			continue
		}
		assignExpression := strings.SplitN(strings.ReplaceAll(line, " ", ""), "=", 2)
		if len(assignExpression) == 2 {
			key := assignExpression[0]
			value := assignExpression[1]

			sections[currentSection][key] = value
		} else {
			return nil, errInvalidSyntax(idx + 1)
		}

	}

	return sections, nil
}

func readFile(file *os.File) ([]string, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return []string{}, err
	}

	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		return []string{}, err
	}

	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	return lines, nil
}

func errInvalidSyntax(lineNumber int) error {
	return errors.New(fmt.Sprintf("Syntax error on line: %d", lineNumber))
}
