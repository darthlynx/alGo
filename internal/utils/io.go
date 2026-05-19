package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetInputLinesFromFile(inputFile string) (_ []string, err error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed opening file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed closing file: %w", cerr)
		}
	}()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

func GetInputLinesFromConsole() []string {
	var inputs []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}
		inputs = append(inputs, text)
	}
	return inputs
}

func WriteOutputToFile(fileName, output string) (err error) {
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed creating file: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed closing file: %w", cerr)
		}
	}()

	if _, err = f.WriteString(output); err != nil {
		return fmt.Errorf("failed writing to file: %w", err)
	}
	return nil
}
