package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type query struct {
	leftEdge  int
	rightEdge int
	count     int
}

// https://acm.timus.ru/problem.aspx?space=1&num=1613&locale=en
func main() {
	inputs := getInputLines()
	// n, err := strconv.Atoi(inputs[0]) // number of cities
	// checkError(err)
	// q, err := strconv.Atoi(inputs[2]) // requests number
	// checkError(err)

	stats := getStatistics(inputs)
	queries := getQueries(inputs)

	for _, q := range queries {
		if v, ok := stats[q.count]; ok {
			if len(v) == 0 {
				fmt.Print(0)
				continue
			}
			floor := lowerBound(v, q.leftEdge)
			if floor > 0 && q.rightEdge >= floor {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		} else {
			fmt.Print(0)
		}

	}

}

// map of stat numbers and their positions in initial array
func getStatistics(inputs []string) map[int][]int {
	statistics := stringToIntArray(inputs[1], " ")
	stats := make(map[int][]int, len(statistics))
	for i, v := range statistics {
		stats[v] = append(stats[v], i+1)
	}
	return stats
}

// binary search to find lower bound
func lowerBound(arr []int, value int) int {
	var l, r, m int
	l = -1
	r = len(arr)
	for l+1 < r {
		m = (l + r) / 2
		if m < len(arr) && arr[m] >= value {
			r = m
		} else {
			l = m
		}
	}
	if r > len(arr)-1 {
		return -1
	} else {
		return arr[r]
	}
}

func getQueries(inputs []string) []query {
	var queries []query
	for i := 3; i < len(inputs); i++ {
		ii := stringToIntArray(inputs[i], " ")
		q := query{
			leftEdge:  ii[0],
			rightEdge: ii[1],
			count:     ii[2],
		}
		queries = append(queries, q)
	}
	return queries
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
	var ii []int
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
