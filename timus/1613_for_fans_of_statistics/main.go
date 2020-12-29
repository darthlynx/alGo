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
	// fmt.Println(stats)
	queries := getQueries(inputs)
	// fmt.Println(queries)

	var res strings.Builder
	// b.Grow(32)
	for _, q := range queries {
		if v, ok := stats[q.count]; ok {
			if len(v) == 0 {
				fmt.Fprintf(&res, "%v", 0)
				continue
			}
			right := UpperBound(v, q.rightEdge)
			if right > 0 && right >= q.leftEdge {
				fmt.Fprintf(&res, "%v", 1)
			} else {
				fmt.Fprintf(&res, "%v", 0)
			}
		} else {
			fmt.Fprintf(&res, "%v", 0)
		}

	}
	fmt.Println(res.String())

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

func UpperBound(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	min := 0
	max := len(arr) - 1
	for min < max {
		mid := (min + max + 1) / 2
		if target >= arr[mid] {
			min = mid
		} else {
			max = mid - 1
		}
	}
	if arr[max] > target {
		return -1
	} else {
		return arr[max]
	}
}

func getQueries(inputs []string) []query {
	queries := []query{}
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
		scanner.Scan()
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
