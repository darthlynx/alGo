package twopointers

import "testing"

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func TestPairSum(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"example 1 - all pairs equal", []int{5, 4, 2, 1}, 6},
		{"example 2 - max at first pair", []int{4, 2, 2, 3}, 7},
		{"example 3 - two nodes max value", []int{1, 100000}, 100001},
		{"all same values", []int{3, 3, 3, 3}, 6},
		{"max in middle pair", []int{1, 5, 3, 2, 4, 1}, 9},
		{"max at last pair", []int{1, 1, 1, 10}, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pairSum(buildList(tt.vals))
			if got != tt.want {
				t.Errorf("pairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
