package utils

import (
	"bufio"
	"log"
	"os"
)

func GetInputLinesFromFile(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("failed closing file: %s", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return lines
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

func WriteOutputToFile(fileName, output string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed closing file: %s", err)
		}
	}()

	if _, err := f.WriteString(output); err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
}
