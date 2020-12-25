package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://acm.timus.ru/problem.aspx?space=1&num=1613&locale=en
func main() {
	// TODO
	inputs := getInputLines()
	n, err := strconv.Atoi(inputs[0]) // number of cities
	fmt.Println(n)
	checkError(err)
	statistics := stringToIntArray(inputs[1], " ")
	q, err := strconv.Atoi(inputs[2]) // requests number
	checkError(err)

	g := make(map[int][]int, q)

	for i, v := range statistics {
		g[v] = append(g[v], i+1)
	}

	fmt.Println(g)

	// for i := 3; i < len(inputs); i++ {
	// 	ii := stringToIntArray(inputs[i], " ")
	// 	graph[ii[len(ii)-1]] = ii[:len(ii)-1]
	// }
	// fmt.Println(n)
	// fmt.Println(statistics)
	// fmt.Println(graph)
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
