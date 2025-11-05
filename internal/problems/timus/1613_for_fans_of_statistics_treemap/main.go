package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
)

type query struct {
	leftEdge  int
	rightEdge int
	count     int
}

type void struct{}

// https://acm.timus.ru/problem.aspx?space=1&num=1613&locale=en
func main() {
	inputs := getInputLines()
	stats := getStatistics(inputs)
	queries := getQueries(inputs)

	for _, q := range queries {
		if v, ok := stats[q.count]; ok {
			floor, _ := v.Floor(q.rightEdge)
			if floor != nil && q.leftEdge <= floor.(int) {
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
func getStatistics(inputs []string) map[int]*treemap.Map {
	statistics := stringToIntArray(inputs[1], " ")
	stats := make(map[int]*treemap.Map, len(statistics))
	for i, v := range statistics {
		if stats[v] == nil {
			stats[v] = treemap.NewWithIntComparator()
		}
		(*stats[v]).Put(i+1, void{})
	}
	return stats
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
