package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	inputs := readInputsFromFile("INPUT.TXT")
	a, err := strconv.Atoi(inputs[0])
	check(err)
	b, err := strconv.Atoi(inputs[1])
	check(err)
	sum := a + b
	fmt.Println(sum)
	writeOutputToFile("OUTPUT.TXT", fmt.Sprintf("%d", sum))
}

func readInputsFromFile(fileName string) []string {
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	inputs := strings.Fields(line)
	fmt.Println(inputs)
	return inputs
}

func writeOutputToFile(fileName, output string) {
	f2, err := os.Create(fileName)
	defer func() {
		err := f2.Close()
		check(err)
	}()
	f2.WriteString(output)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
