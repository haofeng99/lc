package tree

// https://leetcode.cn/problems/recover-binary-search-tree/
// 99. 恢复二叉搜索树
// 需要注意两个节点相邻时, 第二个节点赋值为空的问题
// O(n) O(h)
func recoverTree(root *TreeNode) {
	queue := []*TreeNode{}
	var first, second, pre *TreeNode
	for root != nil || len(queue) > 0 {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		root = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if pre != nil && root.Val < pre.Val {
			second = root // 只要发现逆序对, 直接更新second,
			// 因为假设去除掉其他节点, 当前两节点本身就是前后逆序对的两个节点
			if first == nil {
				first = pre
			}
		}
		pre = root
		root = root.Right
	}

	first.Val, second.Val = second.Val, first.Val
}
