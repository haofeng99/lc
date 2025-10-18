package tree

// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
// https://www.nowcoder.com/practice/947f6eb80d944a84850b0538bf0ec3a5?tpId=13&&tqId=11179&rp=1&ru=/activity/oj&qru=/ta/coding-interviews/question-ranking
// 426-LCR
// 155. 将二叉搜索树转化为排序的双向链表
// 思路: 中序遍历 有序性
// O(n)
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{}
	var ans, pre *TreeNode
	for len(queue) != 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		root = queue[len(queue)-1]
		if ans == nil {
			ans = root
		}
		queue = queue[:len(queue)-1]

		if pre != nil {
			pre.Right = root
			root.Left = pre
		}
		pre = root
		root = root.Right
	}
	return ans
}

// 中序遍历过程dfs O(n)
func treeToDoublyList_dfs(root *TreeNode) *TreeNode {

	var ans, pre *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if ans == nil {
			ans = node
		}
		if pre != nil {
			pre.Right = node
			node.Left = pre
		}
		pre = node
		dfs(node.Right)
	}
	dfs(root)

	return ans
}

// O(1)空间复杂度优化 莫里斯中序遍历
