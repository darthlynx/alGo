package leetcode2196

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/create-binary-tree-from-descriptions/
//
// Time Complexity: O(n),
// Space Complexity: O(n).
func createBinaryTree(descriptions [][]int) *TreeNode {
	hasParent := make(map[int]bool)
	children := make(map[int][2]int)

	for _, desc := range descriptions {
		parent := desc[0]
		child := desc[1]

		hasParent[child] = true
		lr, ok := children[parent]
		if !ok {
			lr = [2]int{-1, -1}
		}
		if desc[2] == 1 {
			lr[0] = child
		} else {
			lr[1] = child
		}
		children[parent] = lr
	}

	rootNode := -1
	for k := range children {
		if !hasParent[k] { // found root node (without parent)
			rootNode = k
			break
		}
	}

	return dfs(rootNode, children)
}

func dfs(nodeVal int, children map[int][2]int) *TreeNode {
	node := TreeNode{
		Val: nodeVal,
	}

	if ch, ok := children[node.Val]; ok {
		if ch[0] != -1 {
			node.Left = dfs(ch[0], children)
		}
		if ch[1] != -1 {
			node.Right = dfs(ch[1], children)
		}
	}

	return &node
}
