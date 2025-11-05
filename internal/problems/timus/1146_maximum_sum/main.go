package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

//4
//0 -2 -7 0
//9 2 -6 2
//-4 1 -4 1
//-1 8 0 -2

// https://acm.timus.ru/problem.aspx?space=1&num=1146
func main() {
	lines := getInputLines()
	arrayDim, _ := strconv.Atoi(lines[0])
	matrix := make([][]int, arrayDim)

	for i := 1; i < len(lines); i++ {
		matrix[i-1] = stringToIntArray(lines[i])
	}

	fmt.Println(maxSum(matrix))
}

// https://www.youtube.com/watch?v=r9W4f5UyMMM&feature=youtu.be
func maxSum(matrix [][]int) int {
	maxSum := MinInt
	numColumns := len(matrix[0])
	for left := 0; left < numColumns; left++ {
		temp := make([]int, numColumns)

		for right := left; right < numColumns; right++ {
			for i := 0; i < len(matrix); i++ {
				temp[i] += matrix[i][right]
			}

			// find max sum in array
			sum := maxSubArraySum(temp)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

// Kadane's algorithm:  https://en.wikipedia.org/wiki/Maximum_subarray_problem#Kadane's_algorithm
func maxSubArraySum(array []int) int {
	maxSum := MinInt
	currentSum := 0
	for _,v := range array {
		currentSum = maxOf(v, currentSum + v)
		maxSum = maxOf(maxSum, currentSum)
	}
	return maxSum
}

func maxOf(a,b int) int {
	if a > b {
		return a
	} else {
		return b
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

func stringToIntArray(str string) []int {
	separator := " "
	arr := strings.Split(str, separator)
	res := make([]int, len(arr))

	for i, v := range arr {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}
