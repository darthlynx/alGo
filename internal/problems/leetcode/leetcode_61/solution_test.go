package leetcode61

import (
	"reflect"
	"testing"
)

func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want []int
	}{
		{"example 1: rotate by 2", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"example 2: rotate by 4", []int{0, 1, 2}, 4, []int{2, 0, 1}},
		{"nil head", nil, 3, nil},
		{"single element", []int{1}, 5, []int{1}},
		{"k equals length (no rotation)", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k is zero", []int{1, 2, 3, 4}, 0, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listToSlice(rotateRight(makeList(tt.vals), tt.k))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
