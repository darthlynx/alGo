package main

import "fmt"

// https://brestprog.by/topics/prefixsums/
func main() {

	matrix := [][]int{
		{0, -2, -7, 0},
		{9, 2, -6, 2},
		{-4, 1, -4, 1},
		{-1, 8, 0, -2},
	}
	fmt.Println("initial matrix")
	printMatrix(matrix)

	pref := setUpPrefixes(matrix)
	fmt.Println("prefix sums")
	printMatrix(pref)

	x1, y1, x2, y2 := 1, 0, 2, 2
	fmt.Printf("Sum for square: x1=%v, y1=%v, x2=%v, y2=%v =>\t", x1, y1, x2, y2)
	fmt.Println(getSumOnRectangle(pref, x1, y1, x2, y2))

}

func setUpPrefixes(matrix [][]int) [][]int {
	pref := make([][]int, len(matrix))
	for i := 0; i < len(pref); i++ {
		pref[i] = make([]int, len(matrix[0]))
	}

	pref[0][0] = matrix[0][0]

	for i := 1; i < len(matrix); i++ {
		pref[i][0] = pref[i-1][0] + matrix[i][0]
	}

	for i := 1; i < len(matrix[0]); i++ {
		pref[0][i] = pref[0][i-1] + matrix[0][i]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			pref[i][j] = pref[i-1][j] + pref[i][j-1] - pref[i-1][j-1] + matrix[i][j]
		}
	}
	return pref
}

func getSumOnRectangle(pref [][]int, x1, y1, x2, y2 int) int {
	result := pref[x2][y2]
	if x1 > 0 {
		result -= pref[x1-1][y2]
	}
	if y1 > 0 {
		result -= pref[x2][y1-1]
	}
	if x1 > 0 && y1 > 0 {
		result += pref[x1-1][y1-1]
	}
	return result
}

func printMatrix(m [][]int) {
	for _, v := range m {
		for _, w := range v {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}