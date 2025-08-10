package tree

// https://leetcode.cn/problems/count-complete-tree-nodes/
// https://cloud.tencent.com/developer/article/1880865
// 222. 完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var deptL func(*TreeNode) int
	deptL = func(node *TreeNode) int {
		ans := 0
		for node != nil {
			ans++
			node = node.Left
		}
		return ans
	}

	var deptR func(*TreeNode) int
	deptR = func(node *TreeNode) int {
		ans := 0
		for node != nil {
			ans++
			node = node.Right
		}
		return ans
	}

	lDept := deptL(root.Left)
	rDept := deptR(root.Right)

	if lDept == rDept {
		return ((1<<lDept)-1)*2 + 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
