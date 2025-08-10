package tree

// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/

// 和通过数组构建的方式不同, 每次需要找中位,比较耗费时间
// 链表为中序遍历结果, 那就直接构建树, 一边中序构建树, 一边填充数值
func sortedListToBST(head *ListNode) *TreeNode {

	// 匿名函数
	nums := func(node *ListNode) int {
		ans := 0
		for node != nil {
			node = node.Next
			ans++
		}
		return ans
	}(head)

	cursor = head

	return dfs_sortedListToBST(0, nums-1)
}

var cursor *ListNode

// 中序遍历反序列化
func dfs_sortedListToBST(left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2

	root := &TreeNode{}
	root.Left = dfs_sortedListToBST(left, mid-1)
	root.Val = cursor.Val
	cursor = cursor.Next
	root.Right = dfs_sortedListToBST(mid+1, right)
	return root
}
