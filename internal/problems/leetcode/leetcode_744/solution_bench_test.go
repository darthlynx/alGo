package leetcode744

import (
	"fmt"
	"testing"
)

func makeLetters(n int) []byte {
	letters := make([]byte, n)
	for i := 0; i < n; i++ {
		letters[i] = byte('a' + (i % 26))
	}
	return letters
}

func BenchmarkNextGreatestLetter(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	for _, n := range sizes {
		letters := makeLetters(n)
		target := byte('a' + byte((n/2)%26))

		b.Run(fmt.Sprintf("size=%d/solution1", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = nextGreatestLetter_solution1(letters, target)
			}
		})

		b.Run(fmt.Sprintf("size=%d/solution2", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = nextGreatestLetter_solution2(letters, target)
			}
		})
	}
}
