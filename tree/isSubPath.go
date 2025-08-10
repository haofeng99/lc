package tree

// https://leetcode.cn/problems/linked-list-in-binary-tree/
// 1367 二叉树中的链表
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	var dfs func(*ListNode, *TreeNode) bool
	dfs = func(list *ListNode, node *TreeNode) bool {
		// list已经匹配完了,说明成功了
		if list == nil {
			return true
		}
		// node已经匹配完了,还没匹配结束,说明匹配失败
		if node == nil {
			return false
		}
		if list.Val != node.Val {
			return false
		}
		return dfs(list.Next, node.Left) || dfs(list.Next, node.Right)
	}
	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
