package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://acm.timus.ru/problem.aspx?space=1&num=1119
func main() {
	inputlines := getInputLines()
	limits := stringToIntArray(inputlines[0], " ")
	n := limits[0] + 1 // increace as we start from 1,1 point
	m := limits[1] + 1 // increace as we start from 1,1 point

	//k := 0
	//if len(inputlines) > 1 {
	//	k, _ = strconv.Atoi(inputlines[1])
	//}

	//for i:=2; i < len(inputlines)-1; i++ {
	//	tmp := stringToIntArray(inputlines[i], " ")

	//}

	distance := 100.0
	diag := math.Sqrt(math.Pow(float64(distance), 2) * 2)

	arr := make([][]float64, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]float64, m)
	}

	for i := 1; i < n; i++ {
		arr[i][0] = arr[i-1][0] + distance
	}
	for i := 1; i < m; i++ {
		arr[0][i] = arr[0][i-1] + distance
	}

	for i := 2; i < len(inputlines)-1; i++ {
		tmp := stringToIntArray(inputlines[i], " ")
		arr[tmp[0]][tmp[1]] = -1
	}

	//arr[1][1] = -1
	//arr[3][2] = -1
	//arr[1][2] = -1

	//print(arr, n, m)

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if arr[i][j] < 0 {
				tmp := arr[i-1][j-1] + diag
				arr[i][j] = math.Min(math.Min(arr[i][j-1]+distance, arr[i-1][j]+distance), tmp)
			} else {
				arr[i][j] = math.Min(arr[i-1][j]+distance, arr[i][j-1]+distance)
			}
		}
	}

	print(arr, n, m)

	res := math.Round(arr[len(arr)-1][len(arr[0])-1])
	fmt.Println("Result:", res)
	writeOutputToFile("output.txt", fmt.Sprintf("%v", res))
}

func print(arr [][]float64, n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()
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

func stringToIntArray(str, sep string) []int {
	arr := strings.Split(str, sep)
	res := make([]int, len(arr))

	for i, v := range arr {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func writeOutputToFile(fileName, output string) {
	f2, _ := os.Create(fileName)
	defer f2.Close()
	f2.WriteString(output)
	//check(err)
}
