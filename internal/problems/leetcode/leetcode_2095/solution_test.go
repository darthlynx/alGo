package leetcode2095

import (
	"reflect"
	"testing"
)

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func toSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestDeleteMiddle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"single node returns nil", []int{1}, nil},
		{"two nodes - delete second", []int{1, 2}, []int{1}},
		{"three nodes - delete middle", []int{1, 2, 3}, []int{1, 3}},
		{"four nodes (example 2)", []int{1, 2, 3, 4}, []int{1, 2, 4}},
		{"seven nodes (example 1)", []int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
		{"five nodes - odd length", []int{1, 2, 3, 4, 5}, []int{1, 2, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSlice(deleteMiddle(buildList(tt.vals)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
