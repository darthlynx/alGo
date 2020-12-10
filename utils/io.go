package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetInputLinesFromFile(inputFile string) []string {
	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines
}

func GetInputLinesFromConsole() []string {
	inputs := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			inputs = append(inputs, text)
		} else {
			break
		}
	}
	return inputs
}

func WriteOutputToFile(fileName, output string) {
	f2, err := os.Create(fileName)
	defer f2.Close()
	f2.WriteString(output)
	if err != nil {
		log.Fatalf("failed writing to a file: %s", err)
	}
}
