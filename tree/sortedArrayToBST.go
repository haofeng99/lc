package tree

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
func sortedArrayToBST(nums []int) *TreeNode {
	return dfs_sortedArrayToBST(nums, 0, len(nums)-1)
}

// 中序遍历反序列化
func dfs_sortedArrayToBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	node := &TreeNode{nums[mid], nil, nil}
	node.Left = dfs_sortedArrayToBST(nums, left, mid-1)
	node.Right = dfs_sortedArrayToBST(nums, mid+1, right)
	return node
}
