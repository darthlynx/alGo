package hareandturtoise

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
//
// Time Complexity: O(n),
// Space Complexity: O(1).
func pairSum(head *ListNode) int {
	hare := head
	turtoise := head
	for hare != nil && hare.Next != nil && turtoise != nil {
		turtoise = turtoise.Next
		hare = hare.Next.Next
	}

	secondHead := turtoise
	secondHead = reverse(secondHead)

	maxPairSum := 0
	for secondHead != nil {
		maxPairSum = max(maxPairSum, head.Val+secondHead.Val)
		head = head.Next
		secondHead = secondHead.Next
	}
	return maxPairSum
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
