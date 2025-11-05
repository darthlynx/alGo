package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://timus.online/problem.aspx?space=1&num=1491&locale=ru
func main() {
	inputs := getInputLines()
	n, err := strconv.Atoi(inputs[0])
	checkError(err)

	nights := make([]int, n+2)

	for i := 1; i < len(inputs); i++ {
		tmp := stringToIntArray(inputs[i], " ")
		nights[tmp[0]] += tmp[2]
		nights[tmp[1]+1] -= tmp[2]
	}

	for i := 1; i < len(nights); i++ {
		nights[i] = nights[i] + nights[i-1]
	}

	for i := 1; i < len(nights)-1; i++ {
		fmt.Printf("%v ", nights[i])
	}
}

func getInputLines() []string {
	inputs := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			inputs = append(inputs, text)
		} else {
			break
		}
	}
	return inputs
}

func stringToIntArray(str, sep string) []int {
	ii := []int{}
	stringArr := strings.Split(str, sep)
	for i := range stringArr {
		tmp, err := strconv.Atoi(stringArr[i])
		checkError(err)
		ii = append(ii, tmp)
	}
	return ii
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("failed writing to a file: %s", err)
	}
}
