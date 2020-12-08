package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	txtlines := getInputLines()
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}

}

func getInputLines() []string {
	file, err := os.Open("input.txt")

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
