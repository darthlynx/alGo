package main

import "fmt"

func main() {

	a := []int{9, 2, -6, 2}
	sums := make([]int, len(a))

	sums[0] = a[0]
	for i := 1; i < len(a); i++ {
		sums[i] = sums[i-1] + a[i]
	}

	fmt.Println("array\t\t", a)
	fmt.Println("prefixes\t", sums)

	fmt.Println(getPref(sums, 1,2))
}

func getPref(sums []int, l, r int) int {
	if l>=0 {
		return sums[r] - sums[l-1]
	} else {
		return sums[r]
	}
}
