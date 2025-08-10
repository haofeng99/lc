package tree

// https://leetcode.cn/problems/binary-search-tree-iterator/solutions/683126/er-cha-sou-suo-shu-die-dai-qi-by-leetcod-4y0y/
// 方法1: 直接在Constructor中序遍历整棵树,并放到队列里面,然后使用一个下标不断遍历队列, 时间复杂度O(n), 空间复杂度O(n)
// 方法2: 按需遍历, 时间复杂度O(n), 单次next()操作O(1) 空间复杂度O(h)
type BSTIterator struct {
	queue []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{queue: []*TreeNode{}}
	for root != nil {
		iterator.queue = append(iterator.queue, root)
		root = root.Left
	}
	return iterator
}

func (this *BSTIterator) Next() int {
	root := this.queue[len(this.queue)-1]
	ans := root.Val
	this.queue = this.queue[:len(this.queue)-1]
	root = root.Right
	for root != nil {
		this.queue = append(this.queue, root)
		root = root.Left
	}
	return ans
}

func (this *BSTIterator) HasNext() bool {
	return len(this.queue) > 0
}
