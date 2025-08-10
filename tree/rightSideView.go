package tree

// https://leetcode.cn/problems/binary-tree-right-side-view/

var view []int

func rightSideView(root *TreeNode) []int {
	view = []int{}
	dfs_rightSideView(root, &view, 0)
	return view
}

func dfs_rightSideView(node *TreeNode, view *[]int, deep int) {
	if node == nil {
		return
	}
	if deep+1 > len(*view) {
		*view = append(*view, -101)
	}
	(*view)[deep] = node.Val
	dfs_rightSideView(node.Left, view, deep+1)
	dfs_rightSideView(node.Right, view, deep+1)
}
