package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var separator = " "

// https://acm.timus.ru/problem.aspx?space=1&num=1280
// topoligical sorting
func main() {

	inputs := getInputLines()

	// read n, m
	ranges := stringToIntArray(inputs[0], separator)
	n := ranges[0] // number of subjects
	m := ranges[1] // limitations number

	nodes := make([][]int, n)

	subjects := stringToIntArray(inputs[len(inputs)-1], separator)

	// fill the nodes and their neighbours
	for i := 1; i <= m; i++ {
		ii := stringToIntArray(inputs[i], separator)
		nodes[ii[0]-1] = append(nodes[ii[0]-1], ii[1])
	}

	// visited nodes
	visited := make([]int, n)

	isCorrectOrder := true
	for _, v := range subjects {
		for _, w := range nodes[v-1] {
			if visited[w-1] != 0 {
				isCorrectOrder = false
				break
			}

		}
		if !isCorrectOrder {
			break
		}
		visited[v-1] = 1
	}

	if isCorrectOrder {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
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
	for i, _ := range stringArr {
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
