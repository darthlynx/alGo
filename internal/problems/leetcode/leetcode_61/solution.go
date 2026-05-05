package leetcode61

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/rotate-list/
//
// Time complexity: O(n).
// Space complexity: O(1).
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	tail := head
	for tail.Next != nil {
		n++
		tail = tail.Next
	}

	shift := k % n
	if shift == 0 {
		return head
	}

	// make list circular
	tail.Next = head

	newTail := head
	for i := 0; i < n-shift-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}
