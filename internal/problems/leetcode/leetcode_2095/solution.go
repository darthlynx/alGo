package leetcode2095

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
//
// Time Complexity: O(n),
// Space Complexity: O(1).
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	prev := head
	hare := head
	turtoise := head

	for hare != nil && hare.Next != nil && turtoise.Next != nil {
		prev = turtoise
		turtoise = turtoise.Next
		hare = hare.Next.Next
	}

	prev.Next = turtoise.Next
	return head
}
