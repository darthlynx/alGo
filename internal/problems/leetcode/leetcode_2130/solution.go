package leetcode2130

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
//
// Time Complexity: O(n),
// Space Complexity: O(1).
func pairSum(head *ListNode) int {
	n := getLen(head)

	secondHead := getSecondHead(head, n/2)
	secondHead = reverse(secondHead)

	maxPairSum := 0
	for secondHead != nil {
		maxPairSum = max(maxPairSum, head.Val+secondHead.Val)
		head = head.Next
		secondHead = secondHead.Next
	}
	return maxPairSum
}

func getLen(head *ListNode) int {
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

func getSecondHead(head *ListNode, half int) *ListNode {
	for i := 0; i < half; i++ {
		head = head.Next
	}
	return head
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
