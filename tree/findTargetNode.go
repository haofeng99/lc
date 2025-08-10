package tree

// LCR 174. 寻找二叉搜索树中的目标节点
// 二叉搜索树中的第K大节点
// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
// 利用二叉搜索树中序遍历有序的特性
func findTargetNode(root *TreeNode, cnt int) int {
	queue := []*TreeNode{}
	vq := []int{}
	for len(queue) > 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		root = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		vq = append(vq, root.Val)
		root = root.Right
	}
	return vq[len(vq)-cnt]
}

// 直接 右 -> 中 -> 左遍历
func findTargetNode_ii(root *TreeNode, cnt int) int {
	ans := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		cnt--
		if cnt == 0 {
			ans = node.Val
		}
		dfs(node.Left)
	}
	dfs(root)
	return ans
}

// 230 二叉搜索树中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	queue := []*TreeNode{}
	for len(queue) > 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		root = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}
