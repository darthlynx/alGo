package leetcode307

// https://leetcode.com/problems/range-sum-query-mutable/
type NumArray struct {
	// segment tree. Each node is a sum of child nodes
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	na := NumArray{
		tree: make([]int, 4*n),
		n:    n,
	}
	if n > 0 {
		na.build(nums, 0, 0, na.n-1)
	}
	return na
}

func (this *NumArray) build(nums []int, node int, left int, right int) {
	if left == right {
		this.tree[node] = nums[left]
		return
	}
	mid := left + (right-left)/2
	this.build(nums, node*2+1, left, mid)
	this.build(nums, node*2+2, mid+1, right)

	this.tree[node] = this.tree[node*2+1] + this.tree[node*2+2]
}

func (this *NumArray) Update(index int, val int) {
	this.update(0, 0, this.n-1, index, val)
}

func (this *NumArray) update(node, left, right, index, value int) {
	if left == right {
		this.tree[node] = value
		return
	}
	mid := left + (right-left)/2
	if index <= mid {
		this.update(node*2+1, left, mid, index, value)
	} else {
		this.update(node*2+2, mid+1, right, index, value)
	}
	this.tree[node] = this.tree[node*2+1] + this.tree[node*2+2]
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.querySum(0, 0, this.n-1, left, right)
}

func (this *NumArray) querySum(node, left, right, ql, qr int) int {
	// no overlap
	if right < ql || left > qr {
		return 0
	}

	// full overlap
	if ql <= left && right <= qr {
		return this.tree[node]
	}

	// partial overlap
	mid := left + (right-left)/2
	leftSum := this.querySum(node*2+1, left, mid, ql, qr)
	rightSum := this.querySum(node*2+2, mid+1, right, ql, qr)

	return leftSum + rightSum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
